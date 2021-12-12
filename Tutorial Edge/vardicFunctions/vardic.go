package main

import "fmt"

func takeaAll(args ...int) {
	fmt.Println(args)
}

func main() {
	takeaAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
