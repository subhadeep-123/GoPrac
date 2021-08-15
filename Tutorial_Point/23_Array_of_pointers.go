package main

import "fmt"

const MAX int = 3

func without_pointer() {
	a := [MAX]int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] - %d\n", i, a[i])
	}
}

func with_pointer() {
	a := [MAX]int{10, 100, 200}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] //assign the address of integer
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}

func with_pointer_matrix_style() {
	arr := [MAX]int{10, 100, 200}
	var i int
	var ptr *int
	ptr = &arr

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr)
		*ptr++
	}
}
func main() {
	// without_pointer()
	// with_pointer()
	with_pointer_matrix_style() // seems like it does not work, need to see afterwards
}
