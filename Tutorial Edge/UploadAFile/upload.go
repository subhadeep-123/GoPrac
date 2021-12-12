package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const DIRNAME string = "temp-images"

func DirMake() {
	_, err := os.Stat(DIRNAME)
	if os.IsNotExist(err) {
		log.Println("Folder does not exists, making one")
		err = os.Mkdir(DIRNAME, 0755)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Error Retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("File Header: %+v\n", handler.Header)

	DirMake()

	tempFile, err := ioutil.TempFile(DIRNAME, "upload-*.png")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tempFile.Close()

	RetrivedFileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	tempFile.Write(RetrivedFileBytes)

	// w.Header().Set(http.StatusAccepted)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
