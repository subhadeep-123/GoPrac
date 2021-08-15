package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("Address of a variable: %x - %d\n", &a, &a)
	fmt.Printf("Address stored in the pointer - %x - %d\n", ip, ip)
	fmt.Printf("Access the value using the pointer - %d\n", *ip)
}
