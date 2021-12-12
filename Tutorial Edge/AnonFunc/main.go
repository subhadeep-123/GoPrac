package main

import "fmt"

func main() {
	yoyoma := func() string {
		return "Giddy Up"
	}
	fmt.Println(yoyoma())
	func() {
		fmt.Println("Checkin another anon function")
	}()
}
