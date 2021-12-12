package main

import "fmt"

func getLimit() func() int {
	limit := 10
	return func() int {
		limit -= 1
		return limit
	}
}
func main() {
	lim := getLimit()
	fmt.Println(lim())
	fmt.Println(lim())
	fmt.Println(lim())
}
