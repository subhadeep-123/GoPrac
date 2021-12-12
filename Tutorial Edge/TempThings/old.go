package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const DIRNAME string = "yoyoma"
const FILENAME string = "giddyUp"

func main() {
	dirInfo, err := os.Stat(DIRNAME)
	if os.IsNotExist(err) {
		log.Println(err.Error())
		os.Mkdir(DIRNAME, os.ModeDir)
	} else {
		fmt.Println(dirInfo.Name())
	}

	file, err := ioutil.TempFile(DIRNAME, FILENAME)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer func() {
		os.Remove(file.Name())
		os.Remove(DIRNAME)
	}()
	fmt.Println(file.Name())
}
