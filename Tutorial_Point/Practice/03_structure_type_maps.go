package main

import "fmt"

type for_maps struct {
	fname string
	lname string
	age   int
	iq    float32
}

func main() {
	var temp_map = map[string]for_maps{
		"Person1": for_maps{"Subhadeep", "Banerjee", 22, 249.9},
		"Person2": for_maps{"Ria", "Gupta", 22, 200.9}}
	fmt.Println(temp_map)
}
