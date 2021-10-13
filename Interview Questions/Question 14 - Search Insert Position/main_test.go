package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		input    []int
		target   int
		expected int
	}{
		{input: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{input: []int{1, 3, 5, 6}, target: 3, expected: 1},
		{input: []int{1, 3, 5, 6}, target: 7, expected: 4},
		{input: []int{1, 3, 5, 6}, target: 0, expected: 0},
		{input: []int{1}, target: 0, expected: 0},
	}

	for _, testCase := range testCases {

		// algorithm
		algorithmResult := SearchInsert(testCase.input, testCase.target)

		// compare
		if algorithmResult != testCase.expected {
			t.Errorf("Wrong, got: %v, want: %v.", algorithmResult, testCase.expected)
		}
	}
}
