package main

import "fmt"

func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums,0, right)
}

func findMaxSubArray(nums []int, left int, right int) int {

	if left == right {
		return nums[left]
	} else {
		mid := (left + right)/2
		leftMax := findMaxSubArray(nums, left, mid)
		rightMax := findMaxSubArray(nums, mid+1, right)
		crossMax := maxCrossing(nums, left, right, mid)
		return max(leftMax, rightMax, crossMax)
	}
}

func maxCrossing(nums []int, left int, right int, mid int) int {

	leftSum := -1 << 31
	rightSum := -1 << 31

	sum := 0

	for i:= mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0

	for i:= mid+1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

func max(values ...int) int {
	maxVal := values[0]
	for _, value := range values {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func main() {
	problem := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	answer := maxSubArray(problem)
	fmt.Println(answer)
}