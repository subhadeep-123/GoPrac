package main

import "fmt"

type person2 struct {
	fname string
	lname string
}
type employee struct {
	person2
	empId int
}

func (p person2) details() {
	fmt.Println(p, " "+"I am a person")
}
func (p employee) details() {
	fmt.Println(p, " "+"I am a employee")
}
func main() {
	p1 := person2{"Subhadeep", "Banerjee"}
	fmt.Println(p1)
	p1.details()
	e1 := employee{person2: p1, empId: 07}
	e1.details()
}
