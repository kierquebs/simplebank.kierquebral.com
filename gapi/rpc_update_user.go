package gapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/kierquebs/simplebank.kierquebral.com/db/sqlc"
	"github.com/kierquebs/simplebank.kierquebral.com/pb"
	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/kierquebs/simplebank.kierquebral.com/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FirstName: sql.NullString{
			String: req.GetFirstName(),
			Valid:  req.FirstName != nil,
		},
		MiddleName: sql.NullString{
			String: req.GetMiddleName(),
			Valid:  req.MiddleName != nil,
		},
		LastName: sql.NullString{
			String: req.GetLastName(),
			Valid:  req.LastName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}
		arg.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	resp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return resp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	if req.FirstName != nil {
		if err := val.ValidateName(req.GetFirstName()); err != nil {
			violations = append(violations, fieldViolation("first_name", err))
		}
	}

	if req.MiddleName != nil {
		if err := val.ValidateName(req.GetMiddleName()); err != nil {
			violations = append(violations, fieldViolation("middle_name", err))
		}
	}

	if req.LastName != nil {
		if err := val.ValidateName(req.GetLastName()); err != nil {
			violations = append(violations, fieldViolation("last_name", err))
		}
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}
