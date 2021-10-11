package main

import "fmt"

func InsertionSort(arr []int) []int {
	// Loop through the array
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		// Loop through the sorted part
		for j := 0; j < i; j++ {
			// Found the right position
			if arr[j] > arr[i] {
				// insertion here
				// shift bigger item to the right
				// arr[j:i+1] = append([]int{},arr[i], arr[j], arr[j+1:i]...)
				for k := i; k > j; k-- {
					arr[k] = arr[k-1]
				}
				arr[j] = temp
				break
			}
		}
	}
	return arr
}

func main() {
	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(InsertionSort(arr))
}
