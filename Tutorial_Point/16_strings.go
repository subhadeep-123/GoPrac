package main

import (
	"fmt"
	"strings"
)

func main() {
	var greeting = "Hello World"
	fmt.Printf("The String is :- %s\n", greeting)

	fmt.Printf("In Hexadecimal Format:- ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}

	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	/*q flag escapes unprintable characters, with + flag it escapses non-ascii
	  characters as well to make output unambigous */
	fmt.Printf("\nQuoted String:- %+q\n", sampleText)

	// Length of string
	fmt.Printf("String is :- %s, and it's length is %d\n", greeting, len(greeting))

	// Concatenated String
	greetings := []string{"Hello", "World!"}
	fmt.Println(strings.Join(greetings, " "))
}
