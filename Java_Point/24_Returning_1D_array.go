package main

import "fmt"

func test() []int {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return arr
}
func showArr(a []int) {
	fmt.Println(a)
}
func main() {
	a := test()
	showArr(a)
}
