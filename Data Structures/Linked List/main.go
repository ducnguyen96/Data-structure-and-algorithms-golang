package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}
// 	bottom up
//	[head]{23}
//	[head]{75}<--[next]{23}
//	[head]{14}<--[next]{75}<--[next]{23}
// 	We use pointer here because we want to change the list
func (l *linkedList) prepend(n *node)  {
	second := l.head
	l.head = n
	l.head.next = second
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
func (l *linkedList) deleteWithValue(value int)  {
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
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{
		data: 23,
	}
	node2 := &node{
		data: 75,
	}
	node3 := &node{
		data: 14,
	}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.printListData()

	myList.deleteWithValue(75)
	myList.printListData()
}