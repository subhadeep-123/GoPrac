package main

import (
	"fmt"
	"path/filepath"
	"regexp"
)

func EmailValidation(input string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	fmt.Printf("Pattern %v\n", re.String())
	fmt.Printf("\nEmail is %v :%v\n", input, re.MatchString(input))
}

func Canonicalization() {
	input := "C:\\Users\\Subhadeep\\Desktop\\12_DEC.lnk"
	resolvedPath, _ := filepath.EvalSymlinks(input)
	fmt.Printf("Input: %v :%v\n", input, resolvedPath)
}
func main() {
	EmailValidation("example<script>alert('Injected!');</script>@domain.com")
	Canonicalization()
}
