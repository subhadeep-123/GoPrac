package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	// Create a slice num with double the capacity of earlier slice
	num := make([]int, len(numbers), cap(numbers)*2)
	printSlice(num)
	copy(num, numbers)
	printSlice(num)
}

func printSlice(x []int) {
	fmt.Printf("Slice - %v, Length - %d, Capcity - %d\n", x, len(x), cap(x))
}
