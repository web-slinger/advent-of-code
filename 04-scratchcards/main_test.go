package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExampleFile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	points, scratchCardNumbers, err := getScratchcardPointsFromFile(wd + "/test.txt")
	if err != nil {
		panic(err)
	}

	require.NoError(t, err)
	assert.Equal(t, 13, points)
	assert.Equal(t, 30, scratchCardNumbers)
}
