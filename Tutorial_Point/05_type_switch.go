package main

import "fmt"

func main(){
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x .%T", i)
	case int:
		fmt.Println("X is int")
	case float64:
		fmt.Println("X is float64")
	case func(int) float64:
		fmt.Println("X is func(int)")
	case bool, string:
		fmt.Println("X is bool or string")
	default:
		fmt.Println("don't know the type")
	}
}