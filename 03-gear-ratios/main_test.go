package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSymbolsForLine(t *testing.T) {
	line := "467..114.."

	partNumbers, symbols := getPotentialPartNumbersForLine(line, 0)
	assert.Len(t, partNumbers, 2)
	assert.Len(t, symbols, 0)
}

func TestGetPotentialPartNumbersWithSymbols(t *testing.T) {

	parts, symbols, err := getPotentialPartNumbersWithSymbols("test.txt")
	require.NoError(t, err)

	validPartNumbers, validPartNumberTotal, invalidPartNumbers, invalidPartNumberTotal := getValidPartNumbers(parts, symbols)

	fmt.Printf("validPartNumberTotal %d\n", validPartNumberTotal)
	fmt.Printf("invalidPartNumberTotal %d\n", invalidPartNumberTotal)

	fmt.Printf("invalidPartNumbers %d\n", len(invalidPartNumbers))
	fmt.Printf("validPartNumbers %d\n", len(validPartNumbers))
	fmt.Printf("potentialPartNumbers %d\n", len(parts))

	for _, part := range validPartNumbers {
		fmt.Printf("%+v\n", part)
	}

	assert.Equal(t, 4361, validPartNumberTotal)
}
