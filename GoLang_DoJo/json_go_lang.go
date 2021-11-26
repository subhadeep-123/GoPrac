package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"fullname"`
	Weapon Weapon `json:"weapon"`
	Level  int    `json:"level"`
}

type Weapon struct {
	Name  string `json:"weapon_name"`
	Level int    `json:"weapon_level"`
}

func main() {
	sFrom := `{"fullname": "Subhadeep", "weapon": {"weapon_name": "Matrix Hand", "weapon_level": 20}, "level": 1}`
	fmt.Println(sFrom)
	fmt.Printf("Type sFrom:- %T\n", sFrom)
	var matrix Ninja
	err := json.Unmarshal([]byte(sFrom), &matrix)
	if err != nil {
		panic(err)
	}
	fmt.Println(matrix)

	sTo, err := json.Marshal(matrix)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", sTo)
	fmt.Printf("%s\n", sTo)
}
