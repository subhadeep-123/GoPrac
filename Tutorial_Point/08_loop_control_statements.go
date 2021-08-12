package main

import "fmt"

func show_break() {
	var a = 10
	fmt.Println("The value of a - ", a)
	for i := 0; i <= 100; i += 2 {
		if i == a {
			fmt.Printf("Matched after %d iteration, result \n", i/2)
			break
		} else {
			fmt.Println("No Match, ", i)
		}
	}
}

func show_continue() {
	a := 20
	for a < 100 {
		if a%2 == 0 {
			a++
			continue
		} else {
			fmt.Println(a)
		}
		a++
	}
}
func show_goto() {
	var a int = 10

LOOP:
	for a < 20 {
		if a%2 == 0 || a == 15 {
			a++
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}

func main() {
	show_break()
	show_continue()
	show_goto()
}
