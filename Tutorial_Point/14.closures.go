package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()

	i := 1
	for i <= 20 {
		fmt.Println(nextNumber())
		i++
	}
}
