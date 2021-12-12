package main

import (
	"fmt"
)

var stage string

func init() {
	fmt.Println("Hello from the other side")
	stage = "Development"
}

func init() {
	fmt.Println("Hello from the other side...2x")
	stage = "Production"
}

func main() {
	fmt.Println("Hey what's up man")
	fmt.Printf("Yo Giddy up, cuz stage is - %s", stage)
}
