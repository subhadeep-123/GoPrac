package main

import (
	"fmt"
	"math"
)

// define an interface
type Shape interface {
	area() float64
}

// define a circle structure
type Circle struct {
	x, y, radius float64
}

// define a rectange structure
type Rectange struct {
	width, height float64
}

// define a method for circle (implementation of Shape.area())
func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for rectange (i)
func main() {

}