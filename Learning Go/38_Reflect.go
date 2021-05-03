package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := 20.05
	fmt.Println(reflect.ValueOf(val))
	fmt.Println(reflect.TypeOf(val))
}
