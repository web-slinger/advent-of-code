package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	// almanac, err := run(wd + "/test.txt")
	// require.NoError(t, err)

	// assert.Len(t, almanac.Seeds, 4)
	// assert.Equal(t, 35, almanac.MinLocation)

	minLocation, err := PartOne(wd + "/test.txt")
	require.NoError(t, err)
	assert.Equal(t, 35, minLocation)

}
