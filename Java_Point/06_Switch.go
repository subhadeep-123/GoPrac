package main

import "fmt"

func main() {
	var val int
	fmt.Println("Choose From the Below Option :- \n1. Odd/Even\n2. Grade Check")
	fmt.Scanln(&val)
	var test int
	fmt.Print("Enter a Number to test :- ")
	fmt.Scanln(&test)
	switch val {
	case 1:
		if test%2 == 0 {
			fmt.Print("Even")
		} else {
			fmt.Print("Odd")
		}
		// fallthrough
	case 2:
		if test >= 50 && test <= 60 {
			fmt.Print("Grade C.")
		} else if test >= 61 && test <= 75 {
			fmt.Print("Grade B.")
		} else if test > 75 && test <= 94 {
			fmt.Print("Grade A.")
		} else if test >= 95 && test <= 100 {
			fmt.Print("Grade A+.")
		} else {
			fmt.Print("Failed")
		}
		// fallthrough
	default:
		fmt.Println("Not a Valid Option.")
	}
}
