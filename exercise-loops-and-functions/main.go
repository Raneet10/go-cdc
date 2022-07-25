package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 1e-18 {
			break
		}
		prev = z
	}

	return z
}

func main() {
	fmt.Println("Square root using Sqrt: ", Sqrt(2))
	fmt.Println("Square root using math.Sqrt: ", math.Sqrt(2))
}
