package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	if x < 0 {
		return math.NaN(), 0
	}
	if x == 0 {
		return 0, 0
	}

	z := x / 2 // Better initial guess
	passes := 0
	maxPasses := 100

	for passes < maxPasses {
		prev := z
		z -= (z*z - x) / (2 * z)
		passes++

		if math.Abs(z-prev) < 1e-10 {
			break
		}
	}

	return z, passes
}

func main() {
	numbers := []float64{2, 9, 1000, 4, -4}

	for _, n := range numbers {
		custom, passes := Sqrt(n)
		if math.IsNaN(custom) {
			fmt.Printf("Sqrt(%v) returned NaN (not a number) as expected.\n\n", n)
			continue
		}
		standard := math.Sqrt(n)
		fmt.Printf("Sqrt(%v):\n  -> Custom:   %v (took %d passes)\n  -> Standard: %v\n\n", n, custom, passes, standard)
	}
}
