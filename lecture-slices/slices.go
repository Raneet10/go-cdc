package lectureslices

import (
	"errors"
	"fmt"
)

func addNumToSlice(n int, arr []int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("Nil Slice")
	}

	for i := range arr {
		arr[i] += n
	}

	return arr, nil
}

func addNumAtEnd(n int, arr []int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("Nil Slice")
	}

	arr = append(arr, n)

	return arr, nil
}

func addNumAtBegg(n int, arr []int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("Nil Slice")
	}

	arr = append([]int{n}, arr...)

	return arr, nil
}

func removeFromPos(arr []int, i int) (int, []int, error) {
	if i < 0 || i >= len(arr) {
		return -1, nil, errors.New("Invalid index")
	}

	if len(arr) == 0 {
		return -1, nil, errors.New("Empty Slice")
	}

	if len(arr) == 1 {
		x := arr[0]
		arr = []int{}

		return x, arr, nil
	}

	x := arr[i]
	arr = append(arr[0:i], arr[i+1:]...)

	fmt.Println("Arr after removing ith elem: ", arr)

	return x, arr, nil
}

// Assuming the elements in a slice are distinct
func mergeWithoutDuplicate(a, b []int) ([]int, error) {
	if a == nil || b == nil {
		return nil, errors.New("Nil Slice")
	}

	dup := make(map[int]struct{})

	for _, v := range a {
		if _, ok := dup[v]; !ok {
			dup[v] = struct{}{}
		}
	}

	for _, v := range b {
		if _, ok := dup[v]; !ok {
			dup[v] = struct{}{}
		} else {
			delete(dup, v)
		}
	}

	c := make([]int, 0, len(dup))

	for k := range dup {
		c = append(c, k)
	}

	return c, nil
}

func removeElemFromSecond(a, b []int) ([]int, error) {
	if a == nil || b == nil {
		return nil, errors.New("Nil Slice")
	}

	dup := make(map[int]struct{})

	for _, v := range b {
		dup[v] = struct{}{}
	}

	for i := 0; i < len(a); i++ {
		if _, ok := dup[a[i]]; ok {
			_, a, _ = removeFromPos(a, i)
			i--
		}
	}

	return a, nil
}

func shiftLeft(a []int, i int) ([]int, error) {
	if i <= 0 || i >= len(a) {
		return nil, errors.New("Invalid index")
	}

	if len(a) < 2 {
		return nil, errors.New("Slice should have at least 2 elements")
	}

	a = append(a[i:], a[:i]...)

	return a, nil
}

func shiftRight(a []int, i int) ([]int, error) {
	if i <= 0 || i >= len(a) {
		return nil, errors.New("Invalid index")
	}

	if len(a) < 2 {
		return nil, errors.New("Slice should have at least 2 elements")
	}

	a = append(a[len(a)-i:], a[:len(a)-i]...)

	return a, nil
}

func copySlice(a []int) ([]int, error) {
	if a == nil {
		return nil, errors.New("Nil Slice")
	}

	b := make([]int, 0, len(a))

	b = append(b, a...)

	return b, nil
}

func swapEvenOdd(a []int) ([]int, error) {
	if len(a) < 2 {
		return nil, errors.New("Slice should have at least 2 elements")
	}

	for i := 0; i < len(a); i += 2 {
		if i == len(a)-1 {
			break
		}

		a[i], a[i+1] = a[i+1], a[i]
	}

	return a, nil
}

func SortInt(a []int, desc bool) []int {
	flag := false
	for i := 0; i < len(a) && !flag; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if desc && a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else {
				flag = true
			}
		}
	}

	return a
}

func SortString(a []string, desc bool) []string {
	flag := false
	for i := 0; i < len(a) && !flag; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if desc && a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else {
				flag = true
			}
		}
	}

	return a
}

func Sort[T int | string](a []T, desc bool) []T {
	flag := false
	for i := 0; i < len(a) && !flag; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if desc && a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			} else {
				flag = true
			}
		}
	}

	return a
}
