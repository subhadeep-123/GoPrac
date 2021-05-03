package main

import "fmt"

func return_anonymous() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	objectFunc := return_anonymous()
	num := 1
	for num <= 10 {
		fmt.Println(objectFunc())
		num++
	}
}
