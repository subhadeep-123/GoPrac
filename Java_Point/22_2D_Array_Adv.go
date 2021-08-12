package main

import (
	"fmt"
)

func addElement() [3][3]int {
	var arr [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var val int
			fmt.Printf("Enter Value for [%d][%d] :- ", i, j)
			fmt.Scanln(&val)
			arr[i][j] = val
		}
	}
	//Returning the 2D Array
	return arr

}

func showArray(arr [3][3]int) {
	fmt.Println(arr)
}

func main() {
	arr := addElement()
	showArray(arr)
}
