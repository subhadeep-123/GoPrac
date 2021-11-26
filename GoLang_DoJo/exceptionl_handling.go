package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic is recovered")
		}
	}()
	const one = 2
	if one != 1 {
		panic("One isn't 1. This isn't good")
	}
}
