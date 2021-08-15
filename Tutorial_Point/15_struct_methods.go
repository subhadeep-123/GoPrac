package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Area of Circle:- %f\n", circle.area())
	fmt.Printf("Perimeter of Circle:- %f\n", circle.perimeter())
}
