package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		arr[i] = i * i
	}

	for j = range arr {
		fmt.Printf("Value of Arr at positing %d - %d\n", j, arr[j])
	}
}
