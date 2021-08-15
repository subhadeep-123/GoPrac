package main

import "fmt"

func main() {

	var arr [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = j
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Printf("\n")
	}
}
