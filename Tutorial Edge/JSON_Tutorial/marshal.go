package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}
type Books struct {
	Title string `json:"title"`
	// Author string `json:"author"`
	Author Author
}

func main() {
	var bookArray []Books

	var i = 0
	for i != 9 {
		entry := Books{Title: "Learning Concurrency in Python", Author: Author{Sales: 3, Age: 25, Developer: true}}
		bookArray = append(bookArray, entry)
		i += 1
	}

	fmt.Println(bookArray)

	// Converting book instance of struct to JSON
	byteArray, err := json.Marshal(bookArray)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(byteArray)
	fmt.Println(string(byteArray))

	fmt.Printf("\n\n\n")

	// This is nice cause it is printed nicely
	byteArray, err = json.MarshalIndent(bookArray, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(byteArray))
}

// Unmarshalling JSON
