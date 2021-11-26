package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello_world_page(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello World")
	case "/ninja":
		fmt.Fprintf(w, "Hello Matrix")
	default:
		fmt.Fprintf(w, "Big Fat Error!!")
	}
	fmt.Printf("Handling the reques with %s Method\n", r.Method)
}

func html_text_page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Did *not* timeout")
}

func main() {
	http.HandleFunc("/", hello_world_page)
	http.HandleFunc("/html", html_text_page)
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}
