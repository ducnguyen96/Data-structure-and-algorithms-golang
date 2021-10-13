package main

import "fmt"

func main() {
	maxs := []int{0, 0, 0}
	nums := []int{2, 7, 9, 3, 1, 6, 8}
	// MAX(n) = max(MAX(n-2), MAX(n-3)) + num[i]
	for index, num := range nums {
		var newMax int = maxs[index+1]
		if maxs[index] > newMax {
			newMax = maxs[index]
		}
		newMax += num
		maxs = append(maxs, newMax)
	}

	fmt.Println(maxs)
}
