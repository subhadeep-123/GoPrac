package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0, errors.New("Math: Negative Numbe Passed to Sqrt")
	}
	return math.Sqrt(val), nil
}
func main() {
	result, err := sqrt(-64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result1, err := sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}
}
