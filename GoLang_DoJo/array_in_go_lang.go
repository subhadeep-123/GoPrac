package main

import "fmt"

func printArray(arr []int) {
	for _, val := range arr {
		fmt.Println(val)
	}
}
func main() {
	var arr [5]int
	fmt.Println(arr)
	arr[2] = 55
	fmt.Println(arr)
	fmt.Println(len(arr))

	b := []int{2, 4, 6, 8}
	fmt.Println(b)

	c := [...]int{1, 3, 5, 7, 9}
	fmt.Println(c)
	printArray(c[:])

	strarr := [...]string{"Subhadeep", "Sweta", "Ria", "Shreya"}
	fmt.Println(strarr)
}
