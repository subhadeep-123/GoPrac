package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Printf("The Type of the String :- %T", x)
}
