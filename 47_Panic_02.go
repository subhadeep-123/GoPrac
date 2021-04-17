package main

import "fmt"

func main() {
	fmt.Println("Calling X From main()")
	x()
	fmt.Println("Returned from X.")
}
func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in X", r)
		}
	}()
	fmt.Println("Executing X..")
	fmt.Println("Calling Y...")
	y(0)
	fmt.Println("Returned Normally From Y.")
}

func y(i int) {
	fmt.Println("Execting Y...")
	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in Y", i)
	fmt.Println("Printing in Y", i)
	y(i + 1)
}
