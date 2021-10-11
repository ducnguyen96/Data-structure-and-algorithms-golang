package main

import "fmt"

func SelectionSort(arr []int) []int {
	l := len(arr)
	// increasing start of array
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		// find the smallest from the start of array to the end
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// swap the found smallest with the start item
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(SelectionSort(arr))
}
