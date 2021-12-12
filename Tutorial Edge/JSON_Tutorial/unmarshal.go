package main

import (
	"encoding/json"
	"fmt"
)

type Src struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Dest struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (d Dest) printInfo() {
	fmt.Println(d.Age, d.Name)
}

func main() {
	matrix := Src{Name: "Subhadeep Banerjee", Age: "22"}
	// val, err := json.Marshal(matrix)
	val, err := json.MarshalIndent(matrix, "", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(val))

	var des Dest
	err = json.Unmarshal(val, &des)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("In Destination")
	des.printInfo()
}
