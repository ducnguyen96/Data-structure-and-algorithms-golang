package main

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items)
	toRemove := s.items[l-1]
	s.items = s.items[:l-1]
	return toRemove
}

//func main() {
//	myStack := Stack{}
//	fmt.Println(myStack)
//	myStack.Push(100)
//	myStack.Push(200)
//	myStack.Push(300)
//
//	fmt.Println(myStack)
//
//	myStack.Pop()
//	fmt.Println(myStack)
//}