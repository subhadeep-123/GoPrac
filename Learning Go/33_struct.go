package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func testArray() {
	var arr = []string{"Subhadeep", "Soumya", "Ria", "Shreya", "Sweta", "Litisha", "Satej", "Shubham", "Suvhradip", "Richard"}
	i := 0
	fmt.Printf("\n\n")
	for i < len(arr) {
		fmt.Println(arr[i])
		i++
	}
}
func main() {
	x := person{firstName: "Subhadeep", lastName: "Banerjee", age: 21}
	fmt.Println(x)
	fmt.Println(x.firstName)

	testArray()
}
