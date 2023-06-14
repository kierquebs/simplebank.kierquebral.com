package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/kierquebs/simplebank.kierquebral.com/db/sqlc"
	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username   string `json:"username" binding:"required,min=6,max=8,alphanum"`
	Password   string `json:"password" binding:"required,min=6,max=8"`
	FirstName  string `json:"first_name" binding:"required,min=3"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	MiddleName        string    `json:"middle_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedpassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedpassword,
		FirstName:      req.FirstName,
		MiddleName:     req.MiddleName,
		LastName:       req.LastName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse{
		Username:          user.Username,
		FirstName:         user.FirstName,
		MiddleName:        user.MiddleName,
		LastName:          user.LastName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}
