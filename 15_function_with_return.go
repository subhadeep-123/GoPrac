package main

import (
	"fmt"
)

func test_func() string {
	var a string
	fmt.Println("Enter Your Name :- ")
	fmt.Scanln(&a)
	return a
}
func main() {
	a := test_func()
	fmt.Printf("You Have Entered - %s", a)
}
