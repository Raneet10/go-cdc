package lecturemaps

import (
	"errors"
	"fmt"

	lectureslices "github.com/raneet10/go-cdc/lecture-slices"
)

func wordOccurence(text string) map[string]int {
	count := make(map[string]int)
	var word []rune

	for _, v := range text {
		if string(v) == " " {
			count[string(word)]++
			word = []rune{}
		} else {
			word = append(word, v)
		}
	}

	fmt.Println("Occurence of each word in the text: ", text)

	for k, v := range count {
		fmt.Println("Word ", k, "occurs %d times", v)
	}

	return count
}

func countIntegers(arr []int) (map[int]int, error) {
	if arr == nil {
		return nil, errors.New("Nil Slice")
	}
	count := make(map[int]int)

	for _, v := range arr {
		count[v]++
	}

	return count, nil
}

func findCommonElements(a, b []int) ([]int, error) {
	if a == nil || b == nil {
		return nil, errors.New("Nil slice(s)")
	}
	common := make(map[int]struct{})
	c := make([]int, 0, len(common))

	for _, v := range a {
		common[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := common[v]; ok {
			c = append(c, v)
		}
	}

	return c, nil
}

func fibonacci(n int) int {
	m := make(map[int]int)

	return fib(n, m)
}

func fib(n int, m map[int]int) int {
	m[0] = 0
	m[1] = 1

	if f, ok := m[n]; ok {
		return f
	}

	a := fib(n-1, m)
	b := fib(n-2, m)

	m[n] = a + b

	return m[n]
}

func sort2DMap(m map[int]map[int]int) {
	fee := make([]int, 0, len(m))

	for k := range m {
		fee = append(fee, k)
	}

	lectureslices.Sort(fee, false)

	for _, v := range fee {
		nonce := make([]int, 0, len(m[v]))

		for k := range m[v] {
			nonce = append(nonce, k)
		}

		lectureslices.Sort(nonce, false)
	}
}
