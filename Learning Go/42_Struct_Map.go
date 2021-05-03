package main

import "fmt"

type Vertex struct {
	lat, lon float64
}

var m = map[string]Vertex{
	"JavaPoint": Vertex{40.68433, -74.39967},
	"SSS-IT":    Vertex{37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
