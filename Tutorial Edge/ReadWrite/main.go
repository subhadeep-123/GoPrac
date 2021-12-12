package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FILENAME string = "matrix.txt"

func CreateAFile() bool {
	data := []byte("Hey this a file ful of data")
	err := ioutil.WriteFile(FILENAME, data, 0777)
	if err != nil {
		fmt.Printf("Shit we got an error - %s\n", err.Error())
		return false
	}
	return true
}

func WritingToanExistentFile(fname string) {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	if _, err = file.WriteString("This is need data yo..\n\n"); err != nil {
		panic(err)
	}
}

func ReadAFile() string {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		fmt.Printf("Shit we got an error - %s\n", err.Error())
	}
	return string(data)
}

func main() {
	if CreateAFile() {
		fmt.Println(ReadAFile())
		fmt.Println("----------------------------------------")
		WritingToanExistentFile(FILENAME)
		fmt.Println(ReadAFile())
		defer os.Remove(FILENAME)
	}
}
