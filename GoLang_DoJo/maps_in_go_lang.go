package main

import "fmt"

func main() {
	var s map[string]string
	fmt.Println(s == nil)
	fmt.Println(s)

	s = map[string]string{}
	fmt.Println(s == nil)
	fmt.Println(s)
}
