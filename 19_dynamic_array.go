package main

import "fmt"

func showArray(arr [10]int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("The array element is - %d", arr[i])
	}
}
func main() {
	var arr [10]int
	for a := 0; a < 10; a++ {
		var val int
		fmt.Println("Enter a value :- ")
		fmt.Scanln(&val)
		arr[a] = val
	}
	showArray(arr)
}
