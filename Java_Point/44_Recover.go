package main

import "fmt"

func testRecover(x, y int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := x / y
	return quotient
}

func main() {
	one := testRecover(10, 0)
	fmt.Println(one)
	two := testRecover(10, 2)
	fmt.Println(two)
}
