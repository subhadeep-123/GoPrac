package main

import "fmt"

func recursive_factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * recursive_factorial(i-1)
}

func recursive_fibo(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return recursive_fibo(i-1) + recursive_fibo(i-2)
}
func main() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d", i, recursive_factorial(i))
	fmt.Println("Fibonacci series")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", recursive_fibo(i))
	}
}
