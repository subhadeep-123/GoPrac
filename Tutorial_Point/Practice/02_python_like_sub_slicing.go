package main

import "fmt"

func main() {
	temp_array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(temp_array[:2])
	fmt.Println(temp_array[2:])
	// fmt.Println(temp_array[:-1])
	fmt.Println(temp_array[:])
	// fmt.Println(temp_array[:11])
}
