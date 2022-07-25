package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	str := strings.Fields(s)

	for _, v := range str {
		count[v] += 1
	}

	return count
}

func main() {
	wc.Test(WordCount)
}
