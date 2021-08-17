package main

import "fmt"

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)

	// SubSlice1 - Starts
	fmt.Printf("numberss[1:4] - %d\n", numbers[1:4])
	fmt.Printf("numberss[:3] - %d\n", numbers[:3])
	fmt.Printf("numberss[4:] - %d\n", numbers[4:])
	// SubSlice1 - End

	// New Slice
	var numbers1 = make([]int, 0, 5)
	printSlice(numbers1)

	// SubSlice2 - Starts
	temp := numbers[:3]
	printSlice(temp)

	temp2 := numbers[2:5]
	printSlice(temp2)
	// SubSlice2 - End
}

func printSlice(x []int) {
	fmt.Printf("Slice - %v, Length - %d, Capcity - %d\n", x, len(x), cap(x))
}
