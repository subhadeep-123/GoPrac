package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile(filename string) {
	stream, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Println(string(stream))
	} else {
		log.Fatal("Error Recieved - %v", err)
	}
}
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		file.WriteString("Writing a small entry to the file.")
	}
	// Checking if the data is written to the file
	readFile(filename)
}
func main() {
	filename := "temp.txt"
	readFile(filename)

	// Writing to a file
	writeFile("newfile.txt")
}
