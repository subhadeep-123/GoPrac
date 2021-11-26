package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello  World")

	var arr = [...]int{1, 2, 3}
	fmt.Println(arr)
	// arr = append(arr, 5, 6, 4)
	// fmt.Println(arr)

	slice := make([]int, 0)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice)

	newSlice := make([]int, 0)
	println(newSlice == nil)
	fmt.Println(newSlice == nil)
}
