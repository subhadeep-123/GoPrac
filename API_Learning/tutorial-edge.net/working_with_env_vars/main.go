package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Environment Variables in go")

	err := os.Setenv("NAME", "Subhadeep Banerjee")
	if err != nil {
		fmt.Println(err)
	}
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Println(name)
	fmt.Println(age)
}
