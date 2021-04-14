package main

import "fmt"

func main() {
	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, true}}
	fmt.Println(s)
}
