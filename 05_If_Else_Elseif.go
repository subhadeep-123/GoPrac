package main

import "fmt"

func main() {
	var val int
	fmt.Print("Enter any Number :- ")
	fmt.Scanf("%d", &val)
	if val%2 == 0 || val > 0 {
		if val >= 10 && val <= 20 {
			fmt.Print("Val is in between 10 and 20")
		} else if val >= 21 && val <= 30 {
			fmt.Print("Val is not in between 21 and 30")
		} else {
			fmt.Print("Val is something else")
		}
	} else {
		fmt.Print("Does not satisfy the initial condition.\n")
	}
}
