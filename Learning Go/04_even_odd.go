package main

import "fmt"

func main() {
	var val int
	fmt.Print("Enter any Number :- ")
	fmt.Scanf("%d", &val)
	if val%2 == 0 {
		fmt.Print("Even.")
	} else {
		fmt.Print("Odd.")
	}
}
