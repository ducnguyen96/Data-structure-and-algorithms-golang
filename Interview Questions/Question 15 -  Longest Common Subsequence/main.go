package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)

	commonSubsequence := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		commonSubsequence[i] = make([]int, l2+1)
	}
	max := 0

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if text1[i] == text2[j] {
				commonSubsequence[i+1][j+1] = commonSubsequence[i][j] + 1
				if commonSubsequence[i+1][j+1] > max {
					max = commonSubsequence[i+1][j+1]
				}
			} else if commonSubsequence[i][j+1] > commonSubsequence[i+1][j] {
				commonSubsequence[i+1][j+1] = commonSubsequence[i][j+1]
			} else {
				commonSubsequence[i+1][j+1] = commonSubsequence[i+1][j]
			}
		}
	}
	return max
}

func main() {
	text1 := "abc"
	text2 := "def"

	result := longestCommonSubsequence(text1, text2)
	fmt.Println("result :", result)
}
