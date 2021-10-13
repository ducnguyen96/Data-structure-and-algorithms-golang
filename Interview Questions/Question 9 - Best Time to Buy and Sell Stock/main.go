package main

import "fmt"

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < l; i++ {
		currentMaxProfit := prices[i] - minPrice
		if currentMaxProfit > maxProfit {
			maxProfit = currentMaxProfit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	max := maxProfit(input)
	fmt.Println("max :", max)
}
