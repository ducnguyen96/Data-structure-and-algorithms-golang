package main

import "fmt"

type maxHeap struct {
	items []int
}

type minHeap struct {
	items []int
}

// insert
func (m *maxHeap) insert(n int)  {
	// insert new item
	m.items = append(m.items, n)

	// heapify
	m.heapifyBottomUp(n)
}

func (m *minHeap) insert(n int)  {
	// insert new item
	m.items = append(m.items, n)

	// heapify
	m.heapifyBottomUp(n)
}

func parent(i int) int {
	return (i-1)/2
}

func left(p int) int {
	return 2*p+1
}

func right(p int) int {
	return 2*p+2
}


func (m *maxHeap) heapifyBottomUp(n int) {
	currentIndex := len(m.items) - 1
	if currentIndex == 0 {
		return
	}
	// heapify bottom up
	// while parent < n ==> swap
	parentIndex := parent(currentIndex)
	for m.items[currentIndex] > m.items[parentIndex] {
		//	swap
		m.swap(parentIndex, currentIndex)
		currentIndex = parentIndex
		parentIndex = parent(parentIndex)
		if parentIndex < 0 {
			return
		}
	}
}

func (m *minHeap) heapifyBottomUp(n int) {
	currentIndex := len(m.items) - 1
	if currentIndex == 0 {
		return
	}
	// heapify bottom up
	// while parent > n ==> swap
	parentIndex := parent(currentIndex)
	for m.items[currentIndex] < m.items[parentIndex] {
		//	swap
		m.swap(parentIndex, currentIndex)
		currentIndex = parentIndex
		parentIndex = parent(parentIndex)
		if parentIndex < 0 {
			return
		}
	}
}

func (m *maxHeap) swap(i1 int, i2 int) {
	m.items[i1], m.items[i2] = m.items[i2], m.items[i1]
}

func (m *minHeap) swap(i1 int, i2 int) {
	m.items[i1], m.items[i2] = m.items[i2], m.items[i1]
}

// extractMax
func (m *maxHeap) extractMax() int {
	// return max
	max := m.items[0]
	l := len(m.items)
	m.items[0] = m.items[l-1]
	m.items = m.items[:l-1]

	// heapify topdown
	m.heapifyTopDown()
	return max
}

// extractMin
func (m *minHeap) extractMin() int {
	// return min
	min := m.items[0]
	l := len(m.items)
	m.items[0] = m.items[l-1]
	m.items = m.items[:l-1]

	// heapify topdown
	m.heapifyTopDown()
	return min
}

func (m *maxHeap) heapifyTopDown()  {
	currentIndex := 0
	lastIndex := len(m.items) - 1

	// if there is only 1 node
	if lastIndex < 1 {
		return
	}

	// currentIndex is on top of the tree
	// We have to constantly swap it with its children until it has no children left

	l, r := left(currentIndex), right(currentIndex)
	childToCompare := 0

	// In the loop we have to assure that index of any node not bigger than lastIndex
	for l <= lastIndex {
		// because the node is added from left to right
		// there is 3 possibilities that can happen
		// - left node is the only child (1)
		// - left node is larger than right node (2)
		// - right node is larger (3)

		// (1) + (2)
		// because node added from left to right
		// and each level have to be fulled before next level added
		// so if 1 node have only 1 child, that child must be at lowest level
		// and also the last index
		if l == lastIndex || m.items[l] > m.items[r] {
			childToCompare = l
		} else {
			// (3)
			childToCompare = r
		}

		if m.items[childToCompare] > m.items[currentIndex] {
			m.swap(currentIndex, childToCompare)
			currentIndex = childToCompare
			l, r = left(currentIndex), right(currentIndex)
		} else {
			return
		}
	}
}

func (m *minHeap) heapifyTopDown()  {
	currentIndex := 0
	lastIndex := len(m.items) - 1

	// if there is only 1 node
	if lastIndex < 1 {
		return
	}

	// currentIndex is on top of the tree
	// We have to constantly swap it with its children until it has no children left

	l, r := left(currentIndex), right(currentIndex)
	childToCompare := 0

	// In the loop we have to assure that index of any node not bigger than lastIndex
	for l <= lastIndex {
		// because the node is added from left to right
		// there is 3 possibilities that can happen
		// - left node is the only child (1)
		// - left node is larger than right node (2)
		// - right node is larger (3)

		// (1) + (2)
		// because node added from left to right
		// and each level have to be fulled before next level added
		// so if 1 node have only 1 child, that child must be at lowest level
		// and also the last index
		if l == lastIndex || m.items[l] < m.items[r] {
			childToCompare = l
		} else {
			// (3)
			childToCompare = r
		}

		if m.items[childToCompare] < m.items[currentIndex] {
			m.swap(currentIndex, childToCompare)
			currentIndex = childToCompare
			l, r = left(currentIndex), right(currentIndex)
		} else {
			return
		}
	}
}

func main() {
	//myHeap := &maxHeap{}
	//items := [...]int{97, 94, 85, 53, 93, 84, 77, 36, 28, 60, 54, 82, 76, 66, 74, 10, 14, 21, 6, 11, 30, 20, 9, 17, 81}
	//for _, item := range items {
	//	myHeap.insert(item)
	//	fmt.Println(myHeap)
	//}
	//for _ = range items {
	//	myHeap.extractMax()
	//	fmt.Println(myHeap)
	//}

	myHeap := &minHeap{}
	items := [...]int{97, 94, 85, 53, 93, 84, 77, 36, 28, 60, 54, 82, 76, 66, 74, 10, 14, 21, 6, 11, 30, 20, 9, 17, 81}
	for _, item := range items {
		myHeap.insert(item)
		fmt.Println(myHeap)
	}
	for _ = range items {
		myHeap.extractMin()
		fmt.Println(myHeap)
	}
}
