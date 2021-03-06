package main

import "fmt"

func MergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	} else {
		// divide
		mid := (l-1)/2
		left := MergeSort(nums[:mid+1])
		right := MergeSort(nums[mid+1:])

		// conquer
		merge := make([]int, l)
		lengthLeft := len(left)
		lengthRight := len(right)
		currentLeftIndex := 0
		currentRightIndex := 0
		currentResultIndex := 0

		// loop through all elements of the two array, compare pair by pair and assign them to the result
		for currentLeftIndex <= lengthLeft - 1 || currentRightIndex <= lengthRight - 1 {
			if currentLeftIndex > lengthLeft - 1{
				merge[currentResultIndex] = right[currentRightIndex]
				currentRightIndex++
				currentResultIndex++
			} else if currentRightIndex > lengthRight - 1{
				merge[currentResultIndex] = left[currentLeftIndex]
				currentLeftIndex++
				currentResultIndex++
			} else if left[currentLeftIndex] < right[currentRightIndex]{
				merge[currentResultIndex] = left[currentLeftIndex]
				currentLeftIndex++
				currentResultIndex++
			} else {
				merge[currentResultIndex] = right[currentRightIndex]
				currentRightIndex++
				currentResultIndex++
			}
		}
		return merge
	}
}

func main() {
	nums := []int{3, 1, 11, 12, 13, 4, 5, 2, 8, 9, 7, 6, 10}
	fmt.Println(MergeSort(nums))
}