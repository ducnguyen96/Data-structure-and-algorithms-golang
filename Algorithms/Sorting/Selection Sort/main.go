package main

import "fmt"

func selectionSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i+1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(selectionSort(arr))
}