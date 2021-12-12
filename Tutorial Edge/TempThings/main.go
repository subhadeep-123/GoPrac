package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	DIRNAME  string = "temp-"
	FILENAME string = "matrix-*.txt"
)

func TempDirMake() string {
	var dirInfo string
	_, err := os.Stat(DIRNAME)
	if os.IsNotExist(err) {
		log.Println("Folder does not exists, making one")
		dirInfo, err = ioutil.TempDir("", DIRNAME)
		if err != nil {
			log.Println(err.Error())
		}

	}
	fmt.Println(dirInfo)
	return dirInfo
}

func TempFileMake(tempDir string) *os.File {
	file, err := ioutil.TempFile(tempDir, FILENAME)
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println(file.Name())
	return file
}

func TempWrite(file *os.File) {
	if _, err := file.Write([]byte("Hey there giddy up man")); err != nil {
		log.Println(err.Error())
	}
}

func TempRead(file *os.File) {
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println(string(data))
}

func main() {
	tempDir := TempDirMake()
	tempFile := TempFileMake(tempDir)

	defer func() {
		os.Remove(tempDir)
		os.Remove(tempFile.Name())
	}()

	TempWrite(tempFile)
	TempRead(tempFile)
}
