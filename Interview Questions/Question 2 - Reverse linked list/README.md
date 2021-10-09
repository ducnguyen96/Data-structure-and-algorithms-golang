```go
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
```