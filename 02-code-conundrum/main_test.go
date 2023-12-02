package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGames(t *testing.T) {

	games, err := getGames("/conundrum.txt")
	require.NoError(t, err)

	assert.Len(t, games, 100)
}
