package main

import (
	"fmt"
	"math"
)

func main() {
	// This are created on the fly, these are declared as function variable
	getSquareRoot := func(x float64) float64 {
		return math.Pow(x, 2)
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("Square Root of %d - %0.2f\n", i, getSquareRoot(float64(i)))
	}
}
