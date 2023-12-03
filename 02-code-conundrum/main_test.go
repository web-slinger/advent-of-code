package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGames(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	gameIdTotals, powerTotals, err := getGameResult(wd + "/conundrum.txt")
	require.NoError(t, err)

	assert.GreaterOrEqual(t, gameIdTotals, 0)
	assert.GreaterOrEqual(t, powerTotals, 0)
}

func TestParseGameLine(t *testing.T) {
	line := "Game 100: 9 green, 2 blue, 12 red; 2 blue, 14 red, 2 green; 14 red, 12 green"

	gameId, gameGrabs, err := parseGameGrabString(line)
	require.NoError(t, err)

	assert.Equal(t, 100, gameId)

	assert.Equal(t, []GameGrab{
		{
			Green: 9,
			Blue:  2,
			Red:   12,
		},
		{
			Blue:  2,
			Red:   14,
			Green: 2,
		},
		{
			Red:   14,
			Green: 12,
		},
	}, gameGrabs)
}

func TestGameExample(t *testing.T) {
	gameIdTotals := 0
	powerTotals := 0

	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	for _, line := range lines {
		gameId, gameGrabs, err := parseGameGrabString(line)
		assert.NoError(t, err)

		power, valid := validateGameGrabPossible(gameId, gameGrabs)
		if valid {
			gameIdTotals += gameId
		}
		powerTotals += power
	}

	assert.Equal(t, 2286, powerTotals)
	assert.Equal(t, 8, gameIdTotals)
}
