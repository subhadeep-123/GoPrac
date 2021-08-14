package main

import "fmt"

func max(num1, num2 int) int {
	fmt.Printf("Value of Num1 - %d, Type - %T\n", num1, num1)
	fmt.Printf("Value of Num2 - %d, Type - %T\n", num2, num2)
	if num1 > num2 {
		return num1
	}
	return num2
}

func return_multiple_values(x, y string) (string, string) {
	return y, x
}

func main() {
	// a := 621
	// b := 596
	// res := max(a, b)
	// fmt.Println(res)

	a, b := return_multiple_values("Subhadeep", "Banerjee")
	fmt.Println(a, b)
}
