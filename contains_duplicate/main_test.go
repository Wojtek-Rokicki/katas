package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	testName       string
	input          []int
	expectedOutput bool
}

var testCases = []TestCase{
	{
		testName:       "Test1",
		input:          []int{1, 2, 3, 1},
		expectedOutput: true,
	},
	{
		testName:       "Test2",
		input:          []int{1, 2, 3, 4},
		expectedOutput: false,
	},
	{
		testName:       "Test3",
		input:          []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		expectedOutput: true,
	},
}

func TestSolutionWithHashMap(t *testing.T) {
	assert := assert.New(t)

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			input := testCase.input
			expectedOutput := testCase.expectedOutput
			output := solutionWithHashMap(input)
			assert.Equal(expectedOutput, output)
		})
	}
}

func TestSolutionWithArray(t *testing.T) {
	assert := assert.New(t)

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			input := testCase.input
			expectedOutput := testCase.expectedOutput
			output := solutionWithArray(input)
			assert.Equal(expectedOutput, output)
		})
	}
}
