package main

import "fmt"

func returnParam(i int) int {
	return i * 1
}
func returnFunction(i int, f func(int) int) int{
	return f(i*i)
}
func main() {
	fmt.Println("Advaned Functions in Go Lang")

	a := returnParam(5)
	fmt.Println(a)

	fmt.Printf("%T\n", returnParam)

	b := returnFunction(5, returnParam)
	fmt.Println(b)
	fmt.Printf("%T\n", returnFunction)

	anonmFunc := func ()  {
		fmt.Println("Anon Function")
	}
	anonmFunc()
}