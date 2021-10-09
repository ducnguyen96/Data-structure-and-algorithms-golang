package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

// 	bottom up
//	[head]{23}
//	[head]{75}<--[next]{23}
//	[head]{14}<--[next]{75}<--[next]{23}
// 	We use pointer here because we want to change the list
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
	if l.tail == nil {
		l.tail = n
	}
}

func (l *linkedList) append(n *node) {
	l.tail.next = n
	l.tail = n
	l.length++
}

//	We use receiver here because we only want to get values from the list
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// [head]{14}<--[next]{75}<--[next]{23}
// suppose we want to delete node has value 75
// we have to update next node address of head to address of node 23
// loop through
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	// we have to keep the previous node to point to next of next node
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		// got to the end of the list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	// delete here
	if previousToDelete.next.data == l.tail.data {
		l.tail = previousToDelete
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *linkedList) reverse()  {
	if l.length < 2 {
		return
	}
	currentNode := l.head
	currentNext := currentNode.next
	for currentNext != nil {
		currentNextNext := currentNext.next

		currentNext.next = currentNode
		currentNode = currentNext
		currentNext = currentNextNext
	}
	l.tail = l.head
	l.head = currentNode
}

func main() {
	myList := &linkedList{}
	node1 := &node{
		data: 23,
	}
	node2 := &node{
		data: 75,
	}
	node3 := &node{
		data: 14,
	}
	node4 := &node{
		data: 20,
	}
	node5 := &node{
		data: 30,
	}
	myList.prepend(node1) // 23
	myList.prepend(node2) // 75 23
	myList.prepend(node3) // 14 75 23
	myList.printListData() // 14 75 23

	myList.deleteWithValue(75) // 14 23
	myList.printListData() // 14 23


	myList.deleteWithValue(23) // 14
	myList.printListData() // 14

	myList.append(node4) // 14 20
	myList.printListData() // 14 20

	myList.prepend(node5) // 30 14 20
	myList.printListData() // 30 14 20


	myList.reverse()
	myList.printListData()
}
