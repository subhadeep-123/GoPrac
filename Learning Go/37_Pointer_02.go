package main

import "fmt"

func main() {
	ptr := new(int)
	fmt.Println("Before - ", ptr, *ptr)
	func(x *int) {
		*x = 10
	}(ptr)
	fmt.Println("After - ", ptr, *ptr)
}
