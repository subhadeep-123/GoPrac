package main

import (
	"fmt"
	"log"
	"net/http"
)

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking to see if Authorized header set...")

		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! we can serve content!")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Homepage")
	fmt.Fprintf(w, "Welcome to the Homepage")
}

func handleRequest() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
