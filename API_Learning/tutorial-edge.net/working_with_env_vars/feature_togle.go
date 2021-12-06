package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("FEATURE_TOGGLE") == "TRUE" {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Fprintf(w, "Exciting New features..")
	} else {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Fprintf(w, "existing boring features")
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}
func main() {
	err := os.Setenv("FEATURE_TOGGLE", "TRUE")
	if err != nil {
		fmt.Errorf("Got Error %s:", err.Error())
	}
	handleRequests()
}
