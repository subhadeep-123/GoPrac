package main

import "fmt"

func main() {
	var arr [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var val int
			fmt.Printf("Enter Value for [%d][%d] :- ", i, j)
			fmt.Scanln(&val)
			arr[i][j] = val
		}
	}
	fmt.Println("Array :- ", arr)
}
