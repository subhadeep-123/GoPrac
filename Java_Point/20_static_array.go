package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	for i, j := range arr {
		fmt.Println(i, j)
	}
}
