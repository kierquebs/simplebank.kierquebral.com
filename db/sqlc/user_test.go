package db

import (
	"context"
	"testing"

	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		FirstName:      util.RandomOwner(),
		MiddleName:     util.RandomOwner(),
		LastName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		HashedPassword: "secret",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.MiddleName, user.MiddleName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreatUser(t *testing.T) {
	createRandomUser(t)
}
