package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {

	owner := RandomOwner()
	require.Len(t, owner, 6)

}

func TestRandomMonet(t *testing.T) {

	money := RandomMoney()
	require.NotZero(t, money)

}

func TestRandomCurrency(t *testing.T) {

	curr := RandomCurrency()
	currencies := []string{"EUR", "USD", "CAD"}

	require.Contains(t, currencies, curr)

}
