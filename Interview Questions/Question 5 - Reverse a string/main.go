package main

import "fmt"

func reverseIteration(w string) string {
	str := ""
	for i := len(w) - 1; i >= 0; i-- {
		str = str + string(w[i])
	}
	return str
}

func reverseRecursion(w string) string {
	wordLength := len(w)
	if wordLength == 1 {
		return w
	}
	return string(w[wordLength-1]) + reverseRecursion(w[:wordLength-1])
}

func main() {
	str := "yoyo mastery"
	fmt.Println(reverseIteration(str))
	fmt.Println(reverseRecursion(str))
}
