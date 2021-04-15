package main

import "fmt"

func return_anonymous_func() func(string) {
	return func(msg string) {
		fmt.Println(msg + " Trying a new things")
	}
}
func say_hello(msg string) {
	fmt.Println(msg)
}
func main() {
	say_hello("Hey")

	//anonymous function declares
	func(msg string) {
		fmt.Println(msg)
	}("Subhadeep is anonymous")

	test_anon := return_anonymous_func()
	test_anon("Subhadeep")
}
