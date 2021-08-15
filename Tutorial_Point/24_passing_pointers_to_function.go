package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("Before swap value of a : %d\n", a)
	fmt.Printf("Before swap value of b : %d\n", b)

	swap(&a, &b)
	fmt.Printf("After swap value of a : %d\n", a)
	fmt.Printf("After swap value of b : %d\n", b)
}
