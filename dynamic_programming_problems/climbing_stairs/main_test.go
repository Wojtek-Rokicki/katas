package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecurrentSolution(t *testing.T) {
	assert := assert.New(t)
	input := 5
	expectedOutput := 8
	output := climbStairsRecurrently(input)
	assert.Equal(expectedOutput, output)
}

func TestDynamicSolution(t *testing.T) {
	assert := assert.New(t)
	input := 5
	expectedOutput := 8
	output := climbStairsDynamic(input)
	assert.Equal(expectedOutput, output)
}
