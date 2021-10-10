package main

import (
	"fmt"
	"math"
)

func fibonacci(i int) int {
	if i < -1 {
		return int(math.Pow(-1, float64(i + 1))) * fibonacci(-i)
	}
	if i == 0 {
		return 0
	}
	if i == 1 || i == -1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	fmt.Println(fibonacci(20))
}
