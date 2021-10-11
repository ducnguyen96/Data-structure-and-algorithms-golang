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

func reverseRune(w string) string {
	runes := []rune(w)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "yoyo mastery"
	fmt.Println(reverseIteration(str))
	fmt.Println(reverseRecursion(str))
	fmt.Println(reverseRune(str))
}
