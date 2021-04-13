package main

import "fmt"

func main() {
	// var a, b, c int
	// a = 10
	// b = 20
	// c = 30
	// fmt.Printf("%T %T %T\n", a, b, c)
	// fmt.Printf("%v %v %v\n", a, b, c)
	// fmt.Print(a, b, c, "\n")
	// fmt.Print("%v %v %v\n", a, b, c) // It does not support any formatting
	// fmt.Println(a, b, c)
	test()
}
func test() {
	var a = 10
	fmt.Print(a, "\n")
	fmt.Printf("%T ", a)
}
