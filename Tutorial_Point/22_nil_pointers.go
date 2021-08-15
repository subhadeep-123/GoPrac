package main

import "fmt"

func main() {
	var ptr *int
	fmt.Printf("The value of ptr is :%x\n", ptr)
	// fmt.Printf("The value that ptr has - %d\n", *ptr) //This is buggy as this does not have any value in the pointer
	if ptr == nil {
		fmt.Printf("The pointer ptr is nil")
	} else {
		fmt.Printf("The pointer is not nil")
	}
}
