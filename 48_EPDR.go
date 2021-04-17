package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hey")
	random()
	checkError(1, 10)
	checkError(10, 5)
}

func random() {
	// defer fmt.Println("0")
	// defer fmt.Println("00") //We can do multiple defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The Function Recovered from panic, Reason - %s\n", r)
		}
	}()
	fmt.Println("1")
	panic("We have paniced")
	fmt.Println("Ubba Lubba Dub Dub")
}

func checkError(a int, b int) {
	total, err := sumOf(a, b)
	if err != nil {
		fmt.Println("There was an error", err)
	} else {
		fmt.Println(total)
	}
}

func sumOf(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("Start Cannot be less than End.")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
