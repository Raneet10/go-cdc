package lecturemaps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordOccurence(t *testing.T) {
	t.Parallel()

	text := "This is a test for maps from a dev in core dev course"

	count := wordOccurence(text)

	require.Equal(t, 2, count["dev"])
}

func TestCountIntegers(t *testing.T) {
	t.Parallel()

	var a []int
	count, err := countIntegers(a)
	require.NotNil(t, err)
	require.Nil(t, count)

	a = []int{1, 1, 2, 3, 4, 4, 4}
	res := 3
	count, err = countIntegers(a)
	require.Nil(t, err)
	require.Equal(t, res, count[4])
}

func TestFindCommonElements(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 5}
	b := []int{5, 4, 1}
	res := []int{1, 5}

	c, err := findCommonElements(a, b)
	require.Nil(t, err)
	require.ElementsMatch(t, res, c)
}

func TestFibonacci(t *testing.T) {
	t.Parallel()

	n := 6
	x := fibonacci(n)
	res := 8

	require.Equal(t, res, x)
}

// TODO: Maybe return the sorted slice of keys and check ?
func TestSort2DMap(t *testing.T) {
	t.Parallel()

	sort2DMap(nil)
}
