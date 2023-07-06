package gapi

import (
	"testing"
	"time"

	db "github.com/kierquebs/simplebank.kierquebral.com/db/sqlc"
	"github.com/kierquebs/simplebank.kierquebral.com/util"
	"github.com/kierquebs/simplebank.kierquebral.com/worker"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
