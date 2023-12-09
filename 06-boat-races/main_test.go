package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRacesFromFile(t *testing.T) {

	wd, err := os.Getwd()
	require.NoError(t, err)

	races, productOfRaceWinWays, err := getRacesFromFile(wd + "/test.txt")
	require.NoError(t, err)
	assert.Len(t, races, 3)
	assert.Equal(t, 288, productOfRaceWinWays)

	races, productOfRaceWinWays, err = getRacesFromFile(wd + "/input.txt")
	require.NoError(t, err)
	assert.Len(t, races, 4)
	assert.Equal(t, 288, productOfRaceWinWays)
}

func TestGetRacesFromFilePart2(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	race, err := getRacesFromFilePart2(wd + "/test.txt")
	require.NoError(t, err)
	assert.Equal(t, 71503, len(race.ButtonHoldsWhichWin))

	race, err = getRacesFromFilePart2(wd + "/input.txt")
	require.NoError(t, err)
	assert.Equal(t, 46561107, len(race.ButtonHoldsWhichWin))
}
