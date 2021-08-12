package main

import "fmt"

func main() {
	var choice int = 4
	switch choice {
	case 1:
		fmt.Printf("This is case - %d", choice)
	case 2:
		fmt.Printf("This is case - %d", choice)
	case 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Printf("This is case - %d", choice)
	default:
		fmt.Printf("Get outta here.")
	}
}
