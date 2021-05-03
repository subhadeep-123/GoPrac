package main

import "fmt"

func main() {
	for a := 1; a <= 3; a++ {
		for b := 3; b >= 1; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}
}
