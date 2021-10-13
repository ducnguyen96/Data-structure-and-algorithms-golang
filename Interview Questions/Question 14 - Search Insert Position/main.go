package main

import "fmt"

func SearchInsert(nums []int, target int) int {
	l := len(nums)
	start := 0
	end := l - 1
	mid := (end - start) / 2

	for end-start > 1 {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
			mid = (end - start) / 2
		} else {
			start = mid
			mid = start + (end-start)/2
		}
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	if target < nums[start] {
		return start
	} else if target > nums[end] {
		return end + 1
	} else {
		return end
	}
}

func main() {
	// nums := []int{1, 3, 5, 6}
	// target := 5
	// result = 2

	// nums := []int{1, 3, 5, 6}
	// target := 3
	// result = 1

	nums := []int{1, 3, 5, 6}
	target := 7
	// result = 4

	// nums := []int{1, 3, 5, 6}
	// target := 0
	// result = 0

	// nums := []int{1}
	// target := 0
	// result = 0

	result := SearchInsert(nums, target)
	fmt.Println(result)
}
