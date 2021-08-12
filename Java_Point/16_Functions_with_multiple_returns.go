package main

import "fmt"

func doStuff(args ...int) (int, float32) {
	addAll := 0
	subAll := 0
	for _, a := range args {
		addAll += a
		subAll -= a
	}
	return addAll, float32(subAll)
}
func main() {
	fmt.Println(doStuff(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
	// doStuff(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
}
