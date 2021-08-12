package main

import (
	"fmt"
)

func test_func() {
	var a string
	fmt.Println("Enter Your Name :- ")
	fmt.Scanln(&a)
	fmt.Printf("You Have Entered - %s", a)
}
func main() {
	test_func()
}
