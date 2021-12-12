package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

const FILENAME string = "user.json"

func main() {
	val, err := os.Open(FILENAME)
	defer val.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteArray, err := ioutil.ReadAll(val)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteArray))

	var users Users
	err = json.Unmarshal(byteArray, &users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("============================================")
	fmt.Println(users)
}
