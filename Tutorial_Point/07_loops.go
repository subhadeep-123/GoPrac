package main

import "fmt"

func for_loops() {
	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}
}

func while_like_for_loops() {
	a := 1
	b := 10
	for a <= b {
		fmt.Printf("Value - %d\n", a)
		a++
	}
}

func iterative_for_loops() {
	numbers := [10]uint64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	// also acts like enumerate, like in python
	for i, x := range numbers {
		fmt.Printf("Array values, %d | %d\n", i, x)
	}
}

func nested_for_loops() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
func main() {
	var num int = 5
	if num == 1 {
		for_loops()
	} else if num == 2 {
		while_like_for_loops()
	} else if num == 3 {
		iterative_for_loops()
	} else if num == 4 {
		nested_for_loops()
	} else {
		fmt.Println("Choose a correct option")
	}
}
