package main

import "fmt"

func main() {
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance[4] = 80.0
	for a := range balance {
		fmt.Println(balance[a])
	}
	fmt.Printf("The Length of the array is - %d\n", len(balance))
}
