package main

import "fmt"

const LEN = 10

func slice_it_up(arr []int, start, end int) {
	fmt.Println(arr[start:end])
}
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, val := range array {
		fmt.Println(idx, val)
	}
	slice_it_up(array, 2, 5)

	var slice_able = make([]int, 5, 20)
	slice_it_up(slice_able, 2, 5)
}
// the [10] nees to be same for the declared also for the argument