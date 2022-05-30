package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordsList := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, v := range wordsList {
		if _, ok := wordsMap[v]; !ok {
			wordsMap[v] = 0
		}
		wordsMap[v] = wordsMap[v] + 1
	}
	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
