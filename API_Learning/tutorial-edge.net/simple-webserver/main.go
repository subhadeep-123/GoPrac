package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi")
	})
	http.HandleFunc("/echo", echoString)
	http.HandleFunc("/increment", incrementCounter)

	log.Fatal(http.ListenAndServe("", nil))
}
