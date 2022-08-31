package lectureslices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddNumToSlice(t *testing.T) {
	// Nil Slice
	var arr []int
	n := 5

	arr, err := addNumToSlice(n, arr)
	require.NotNil(t, err)
	require.Nil(t, arr)

	// Normal Slice
	arr = []int{2, 3, 5}
	res := []int{7, 8, 10}

	arr, err = addNumToSlice(n, arr)
	require.Nil(t, err)
	require.Equal(t, arr, res)
}

func TestAddNumAtEnd(t *testing.T) {
	// Nil Slice
	var arr []int
	n := 5

	arr, err := addNumAtEnd(n, arr)
	require.NotNil(t, err)
	require.Nil(t, arr)

	// Normal Slice
	arr = []int{2, 3}
	res := []int{2, 3, 5}

	arr, err = addNumAtEnd(n, arr)
	require.Nil(t, err)
	require.Equal(t, arr, res)
}

func TestAddNumAtBegg(t *testing.T) {
	// Nil Slice
	var arr []int
	n := 5

	arr, err := addNumAtBegg(n, arr)
	require.NotNil(t, err)
	require.Nil(t, arr)

	// Normal Slice
	arr = []int{2, 3}
	res := []int{5, 2, 3}

	arr, err = addNumAtBegg(n, arr)
	require.Nil(t, err)
	require.Equal(t, arr, res)
}

func TestRemoveFromPos(t *testing.T) {
	// Empty array
	arr := []int{}
	i := 5
	x, arr, err := removeFromPos(arr, i)
	require.Equal(t, x, -1)
	require.Nil(t, arr)
	require.NotNil(t, err)

	// Invalid index
	arr = []int{2, 3, 5}
	i = -1

	x, arr, err = removeFromPos(arr, i)
	require.Equal(t, x, -1)
	require.Nil(t, arr)
	require.NotNil(t, err)

	// Normal Case
	arr = []int{2, 3, 5}
	i = 1
	res := []int{2, 5}

	x, arr, err = removeFromPos(arr, i)
	require.Nil(t, err)
	require.Equal(t, arr, res)
	require.Equal(t, 3, x)
}

func TestMergeWithoutDuplicate(t *testing.T) {
	var a, b []int
	c, err := mergeWithoutDuplicate(a, b)
	require.Nil(t, c)
	require.NotNil(t, err)

	a = []int{1, 2, 5}
	b = []int{5, 3, 2}
	res := []int{1, 3}
	c, err = mergeWithoutDuplicate(a, b)
	require.Nil(t, err)
	require.Equal(t, res, c)
}

func TestRemoveElemFromSecond(t *testing.T) {
	var a, b []int
	a, err := removeElemFromSecond(a, b)
	require.Nil(t, a)
	require.NotNil(t, err)

	a = []int{1, 2, 5}
	b = []int{5, 3, 2}
	res := []int{1}
	a, err = removeElemFromSecond(a, b)
	require.Nil(t, err)
	require.Equal(t, res, a)
}

func TestShiftLeft(t *testing.T) {
	// Nil Slice
	var a []int
	i := 2
	a, err := shiftLeft(a, i)
	require.NotNil(t, err)
	require.Nil(t, a)

	// Invalid index
	a = []int{1, 2, 5, 6}
	i = 4
	a, err = shiftLeft(a, i)
	require.NotNil(t, err)
	require.Nil(t, a)

	// Normal case
	a = []int{1, 2, 5, 6}
	i = 2
	res := []int{5, 6, 1, 2}
	a, err = shiftLeft(a, i)
	require.Nil(t, err)
	require.Equal(t, res, a)
}

func TestShiftRight(t *testing.T) {
	// Nil Slice
	var a []int
	i := 2
	a, err := shiftLeft(a, i)
	require.NotNil(t, err)
	require.Nil(t, a)

	// Invalid index
	a = []int{1, 2, 5, 6}
	i = 4
	a, err = shiftLeft(a, i)
	require.NotNil(t, err)
	require.Nil(t, a)

	// Normal case
	a = []int{1, 2, 5, 6}
	i = 1
	res := []int{2, 5, 6, 1}
	a, err = shiftLeft(a, i)
	require.Nil(t, err)
	require.Equal(t, res, a)
}

func TestCopySlice(t *testing.T) {
	var a []int
	b, err := copySlice(a)
	require.NotNil(t, err)
	require.Nil(t, b)

	a = []int{1, 2, 5}
	b, err = copySlice(a)
	require.Nil(t, err)
	require.Equal(t, b, a)
}

func TestSwapEvenOdd(t *testing.T) {
	// Not enough elements
	a := []int{1}
	a, err := swapEvenOdd(a)
	require.NotNil(t, err)
	require.Nil(t, a)

	// Odd elements
	a = []int{1, 2, 5}
	res := []int{2, 1, 5}
	a, err = swapEvenOdd(a)
	require.Nil(t, err)
	require.Equal(t, res, a)

	// Even elements
	a = []int{1, 2, 5, 6}
	res = []int{2, 1, 6, 5}
	a, err = swapEvenOdd(a)
	require.Nil(t, err)
	require.Equal(t, res, a)
}

func TestSortInt(t *testing.T) {
	a := []int{1, 7, 2, 9, 0}
	desc := false
	res := []int{0, 1, 2, 7, 9}
	b := Sort(a, desc)
	require.Equal(t, res, b)

	desc = true
	res = []int{9, 7, 2, 1, 0}
	b = Sort(a, desc)
	require.Equal(t, res, b)

	s := []string{"c", "o", "r", "e"}
	desc = false
	resStr := []string{"c", "e", "o", "r"}
	s = SortString(s, desc)
	require.Equal(t, resStr, s)

	desc = true
	resStr = []string{"r", "o", "e", "c"}
	s = SortString(s, desc)
	require.Equal(t, resStr, s)
}
