package main

import "fmt"

func PrintArray_1D(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func PrintArray_2D(arr [5][5]int) {
	for i := 0; i < 5; i++ {
		{
			for j := 0; j < 5; j++ {
				fmt.Printf("%d ", arr[i][j])
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	var arr_1d [10]int
	var arr_2d [5][5]int
	var i, j int

	for i = 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			arr_2d[i][j] = j
		}
	}
	PrintArray_2D(arr_2d)

	for j = 0; j < 10; j++ {
		arr_1d[i] = j * j
	}
	PrintArray_1D(arr_1d[:])
}
