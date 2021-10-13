package main

import "fmt"

func climbStairs(n int) int {
	ways := []int{1, 2}
	for i := 2; i < n; i++ {
		ways = append(ways, ways[i-1]+ways[i-2])
	}
	return ways[n-1]
}

func main() {
	n := 10
	ways := climbStairs(n)
	fmt.Println("ways :", ways)
}
