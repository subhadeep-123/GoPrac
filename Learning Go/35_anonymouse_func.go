package main

import "fmt"

func demo() func(name string) {
	return func(name string) {
		fmt.Println(name, ""+"This is an anonymous function")
	}
}
func main() {
	a := demo()
	a("Subhadeep")
}
