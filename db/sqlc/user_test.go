package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		FirstName:      util.RandomOwner(),
		MiddleName:     util.RandomOwner(),
		LastName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
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

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.MiddleName, user2.MiddleName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFirstName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		FirstName: sql.NullString{
			String: newFirstName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyMiddleName(t *testing.T) {
	oldUser := createRandomUser(t)

	newMiddleName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		MiddleName: sql.NullString{
			String: newMiddleName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, newMiddleName, updatedUser.MiddleName)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyLastName(t *testing.T) {
	oldUser := createRandomUser(t)

	newLastName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		LastName: sql.NullString{
			String: newLastName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	newMiddleName := util.RandomOwner()
	newLastName := util.RandomOwner()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		FirstName: sql.NullString{
			String: newFirstName,
			Valid:  true,
		},
		MiddleName: sql.NullString{
			String: newMiddleName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: newLastName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.NotEqual(t, oldUser.MiddleName, updatedUser.MiddleName)
	require.Equal(t, newMiddleName, updatedUser.MiddleName)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
}
