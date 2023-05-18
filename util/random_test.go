package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()

	require.Len(t, owner, 6)
}
