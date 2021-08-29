package main

import (
	"fmt"
	"net/http"
)

func Handler_1(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello World\n")
}

func Handler_2(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello World This is Matrix\n")
}

func main() {
	http.HandleFunc("/", Handler_1)
	http.HandleFunc("/matrix", Handler_2)
	http.ListenAndServe(":8080", nil)
}
