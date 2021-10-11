package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		arr   []int
		k     int
		exist bool
	}{
		{arr: []int{9, 4, 20, 1, 6, 15, 170}, k: 15, exist: true},
		{arr: []int{9, 4, 20, 1, 6, 15, 170}, k: 6, exist: true},
		{arr: []int{9, 4, 20, 1, 6, 15, 170}, k: 10, exist: false},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.BinarySearch(testCase.k)

		if testCase.exist != algorithmResult {
			t.Errorf("Wrong, got: %v, want: %v.", testCase.exist, algorithmResult)
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	testCases := []struct {
		arr    []int
		search []int
	}{
		{arr: []int{9, 4, 20, 1, 6, 15, 170}, search: []int{9, 4, 20, 1, 6, 15, 170}},
		{arr: []int{9, 6, 12, 1, 4, 34, 45}, search: []int{9, 6, 12, 1, 34, 4, 45}},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.BreadthFirstSearch()

		for index, _ := range testCase.search {
			if testCase.search[index] != algorithmResult[index] {
				t.Errorf("Wrong, got: %v, want: %v.", algorithmResult[index], testCase.search[index])
			}
		}
	}
}

func TestBreadthFirstSearchRecursive(t *testing.T) {
	testCases := []struct {
		arr    []int
		search []int
	}{
		{arr: []int{9, 4, 20, 1, 6, 15, 170}, search: []int{9, 4, 20, 1, 6, 15, 170}},
		{arr: []int{9, 6, 12, 1, 4, 34, 45}, search: []int{9, 6, 12, 1, 34, 4, 45}},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.BreadthFirstSearchRecursive([]*Node{tree}, []int{})

		for index := range testCase.search {
			if testCase.search[index] != algorithmResult[index] {
				t.Errorf("Wrong, got: %v, want: %v.", algorithmResult[index], testCase.search[index])
			}
		}
	}
}

func TestDepthFirstSearchInOrder(t *testing.T) {
	testCases := []struct {
		arr    []int
		search []int
	}{
		{arr: []int{9, 4, 6, 20, 170, 15, 1}, search: []int{1, 4, 6, 9, 15, 20, 170}},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.DepthFirstSearchInOrder()

		for index := range testCase.search {
			if testCase.search[index] != algorithmResult[index] {
				t.Errorf("Wrong, got: %v, want: %v.", algorithmResult[index], testCase.search[index])
			}
		}
	}
}

func TestDepthFirstSearchPreOrder(t *testing.T) {
	testCases := []struct {
		arr    []int
		search []int
	}{
		{arr: []int{9, 4, 6, 20, 170, 15, 1}, search: []int{9, 4, 1, 6, 20, 15, 170}},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.DepthFirstSearchPreOrder()

		for index := range testCase.search {
			if testCase.search[index] != algorithmResult[index] {
				t.Errorf("Wrong, got: %v, want: %v.", algorithmResult[index], testCase.search[index])
			}
		}
	}
}

func TestDepthFirstSearchPostOrder(t *testing.T) {
	testCases := []struct {
		arr    []int
		search []int
	}{
		{arr: []int{9, 4, 6, 20, 170, 15, 1}, search: []int{1, 6, 4, 15, 170, 20, 9}},
	}

	for _, testCase := range testCases {
		tree := &Node{Key: testCase.arr[0]}
		for _, node := range testCase.arr {
			tree.Insert(node)
		}

		algorithmResult := tree.DepthFirstSearchPostOrder()

		for index := range testCase.search {
			if testCase.search[index] != algorithmResult[index] {
				t.Errorf("Wrong, got: %v, want: %v.", algorithmResult[index], testCase.search[index])
			}
		}
	}
}
