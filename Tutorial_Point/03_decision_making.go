package main

import "fmt"

func main() {
	a := 20
	if a > 0 && a <= 10 {
		fmt.Printf("Huraay, a - %d", a)
	} else if a == 15 {
		fmt.Printf("Damn a is equals to %d", a)
	} else if a == 25 {
		fmt.Printf("Damn a is again something else - %d", a)
	} else {
		if a > 25 {
			fmt.Printf("Thank god a is something - %d", a)
		} else {
			fmt.Printf("I am outta here.")
		}
	}
}
