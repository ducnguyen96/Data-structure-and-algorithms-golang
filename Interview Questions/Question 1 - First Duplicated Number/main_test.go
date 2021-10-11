package main

import "testing"

func TestFirstDuplicate(t *testing.T) {
	testCases := []struct {
		arr []int
		dup int
	}{
		{arr: []int{2, 5, 1, 2, 3, 5, 1, 2, 4}, dup: 2},
		{arr: []int{2, 1, 1, 2, 3, 5, 1, 2, 4}, dup: 1},
		{arr: []int{5, 5, 1, 2, 3, 5, 1, 2, 4}, dup: 5},
	}

	for _, testCase := range testCases {
		algorithmResult, _ := FirstDuplicate(testCase.arr)

		if testCase.dup != algorithmResult {
			t.Errorf("Wrong, got: %v, want: %v.", algorithmResult, testCase.dup)
		}
	}
}
