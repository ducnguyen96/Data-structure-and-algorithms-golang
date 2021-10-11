package main

import "fmt"

func FirstDuplicate(arr []int) (int, bool) {
	myMap := map[int]bool{}
	for _, element := range arr {
		if myMap[element] == true {
			return element, true
		} else {
			myMap[element] = true
		}
	}
	return 0, false
}

func main() {
	// tests
	test1 := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	dup, ok := FirstDuplicate(test1)
	if ok {
		fmt.Println("test 1 dup :", dup)
	} else {
		fmt.Println("test 1 no dup")
	}
	test2 := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	dup, ok = FirstDuplicate(test2)
	if ok {
		fmt.Println("test 2 dup :", dup)
	} else {
		fmt.Println("test 2 no dup")
	}
	test3 := []int{2, 3, 4, 5}
	dup, ok = FirstDuplicate(test3)
	if ok {
		fmt.Println("test 3 dup :", dup)
	} else {
		fmt.Println("test 3 no dup")
	}
}
