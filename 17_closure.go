package main

import "fmt"

func main() {
	num := 2
	result := func() int {
		num *= 2
		return num
	}
	for a := 1; a <= 10; a++ {
		fmt.Printf("The Squared Number is - %v\n", result())
	}

	val := 0
	result_2 := func() int {
		val += 1
		return val
	}
	n := 1
	for n < 11 {
		fmt.Printf("The Added Number is - %v\n", result_2())
		n++
	}
}
