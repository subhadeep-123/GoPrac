package main

import "fmt"

func printPointer(x *int) {
	fmt.Println("The Value is - ", *x)
}
func main() {
	x := 10
	printPointer(&x)
	fmt.Println(x, &x)
}
