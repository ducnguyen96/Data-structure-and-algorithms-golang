package main

import (
	"testing"
)

// [2,1,3]
// [5,1,4,null,null,3,6]
// [32,26,47,19,null,null,56,null,27, 55] v
// [9, 4, 20, 1, 6, 15, 170]
// [1, 1]
// [0,nil,-1]
// [3,1,5,0,2,4,6,null,null,null,3]
// [5,4,6,null,null,3,7]

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		nodeValues []interface{}
		isValid    bool
	}{
		{nodeValues: []interface{}{2, 1, 3}, isValid: true},
		{nodeValues: []interface{}{5, 1, 4, nil, nil, 3, 6}, isValid: false},
		{nodeValues: []interface{}{32, 26, 47, 19, nil, nil, 56, nil, 27, 55}, isValid: false},
		{nodeValues: []interface{}{9, 4, 20, 1, 6, 15, 170}, isValid: true},
		{nodeValues: []interface{}{1, 1}, isValid: false},
		{nodeValues: []interface{}{0, nil, -1}, isValid: false},
		{nodeValues: []interface{}{3, 1, 5, 0, 2, 4, 6, nil, nil, nil, 3}, isValid: false},
		{nodeValues: []interface{}{5, 4, 6, nil, nil, 3, 7}, isValid: false},
	}

	for _, testCase := range testCases {
		// build tree
		root := &TreeNode{Val: testCase.nodeValues[0].(int)}
		values := testCase.nodeValues[1:]
		nodes := make([]*int, len(values))
		var index *int = new(int)
		*index = 0
		for i, v := range values {
			var addr *int
			if v != nil {
				addr = new(int)
				*addr = v.(int)
			}
			nodes[i] = addr
		}

		NotValidateInsert([]*TreeNode{root}, nodes, index)

		// algorithm
		algorithmResult := IsValidBST(root)

		// compare
		if algorithmResult != testCase.isValid {
			t.Errorf("Wrong, got: %v, want: %v.", algorithmResult, testCase.isValid)
		}
	}
}
