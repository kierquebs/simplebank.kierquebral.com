package gapi

import (
	db "github.com/kierquebs/simplebank.kierquebral.com/db/sqlc"
	"github.com/kierquebs/simplebank.kierquebral.com/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FirstName:         user.FirstName,
		MiddleName:        user.MiddleName,
		LastName:          user.LastName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertAccount(account db.Account) *pb.Account {
	return &pb.Account{
		Id:        account.ID,
		Owner:     account.Owner,
		Balance:   account.Balance,
		Currency:  account.Currency,
		CreatedAt: timestamppb.New(account.CreatedAt),
	}
}
