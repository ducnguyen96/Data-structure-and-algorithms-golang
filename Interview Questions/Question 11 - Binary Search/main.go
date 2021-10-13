package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l := len(nums)
	start := 0
	mid := l / 2
	end := l - 1

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	// we keep finding until there is no midle to compre.
	for end-start > 1 {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
			mid = (end - start) / 2
		} else {
			start = mid
			mid = mid + (end-start)/2
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12, 15}
	find := 12
	// find := 5
	result := search(nums, find)
	fmt.Println("result :", result)
}
