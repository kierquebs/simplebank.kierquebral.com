package gapi

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	db "github.com/kierquebs/simplebank.kierquebral.com/db/sqlc"
	"github.com/kierquebs/simplebank.kierquebral.com/pb"
	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/kierquebs/simplebank.kierquebral.com/val"
	"github.com/kierquebs/simplebank.kierquebral.com/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedpassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedpassword,
			FirstName:      req.GetFirstName(),
			MiddleName:     req.GetMiddleName(),
			LastName:       req.GetLastName(),
			Email:          req.GetEmail(),
		},
		AfterCreate: func(user db.User) error {

			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}
	return resp, status.Errorf(codes.OK, "User successfully created")
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("first_name", err))
	}

	if err := val.ValidateName(req.GetMiddleName()); err != nil {
		violations = append(violations, fieldViolation("middle_name", err))
	}

	if err := val.ValidateName(req.GetLastName()); err != nil {
		violations = append(violations, fieldViolation("last_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
