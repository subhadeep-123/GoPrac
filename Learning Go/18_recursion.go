package main

import "fmt"

func fact(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * fact(num-1)
}
func main() {
	fmt.Println("Enter a no :- ")
	var a int
	fmt.Scanln(&a)
	fmt.Printf("The Factorail of %d is %d", a, fact(a))
}
