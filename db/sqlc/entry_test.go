package db

import (
	"context"
	"testing"

	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}
