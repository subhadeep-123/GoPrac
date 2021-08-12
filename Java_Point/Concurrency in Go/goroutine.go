package main

import (
	"fmt"
	"time"
)
func count(thing string) {
	for i := 1; ; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	go count("sheep")
	count("fish")
}
