// go test -v
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	testName       string
	s1             string
	s2             string
	expectedOutput bool
}

var testCases = []TestCase{
	{
		testName:       "Test1",
		s1:             "anagram",
		s2:             "nagaram",
		expectedOutput: true,
	},
	{
		testName:       "Test2",
		s1:             "rat",
		s2:             "car",
		expectedOutput: false,
	},
}

func TestSolutionWithCounterMap(t *testing.T) {
	assert := assert.New(t)

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			s1 := testCase.s1
			s2 := testCase.s2
			expectedOutput := testCase.expectedOutput
			output := solutionWithCounterMap(s1, s2)
			assert.Equal(expectedOutput, output)
		})
	}
}

func TestSolutionWithCounterArray(t *testing.T) {
	assert := assert.New(t)

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			s1 := testCase.s1
			s2 := testCase.s2
			expectedOutput := testCase.expectedOutput
			output := solutionWithCounterArray(s1, s2)
			assert.Equal(expectedOutput, output)
		})
	}
}
