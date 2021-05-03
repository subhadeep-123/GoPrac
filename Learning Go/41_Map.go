package main

import "fmt"

var x = map[string]string{
	"First Name": "Subhadeep",
	"Last Name":  "Banerjee",
	"Age":        "21",
	"Email-id":   "subhadeep762@gmail.com",
	"Hobby":      "Programming"}

func Update() {
	x["Friends"] = "Few"
}
func Delete() {
	delete(x, "Friends")
}
func Interate() {
	for val, Ok := range x {
		fmt.Println(val, Ok)
	}
}
func main() {
	fmt.Println(x)
	Update()
	fmt.Println(x)
	Delete()
	fmt.Println(x)
	Interate()
	//This is what Retrive is used as
	v, ok := x["Hobby"]
	fmt.Println(v, ok)
}
