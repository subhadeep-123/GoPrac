package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSliceDetails(numbers)

	// Nil slice
	var tempSlice []int
	printSliceDetails(tempSlice)
}
func printSliceDetails(x []int) {
	if x == nil {
		fmt.Println("The Slice is nil")
	}
	fmt.Printf("The Length, Capacity and the Slice is %d, %d, and %v\n", len(x), cap(x), x)
}
