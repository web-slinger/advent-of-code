package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolsForLine(t *testing.T) {
	line := "467..114.."

	partNumbers, symbols := getPotentialPartNumbersForLine(line, 0)
	assert.Len(t, partNumbers, 2)
	assert.Len(t, symbols, 0)
}
