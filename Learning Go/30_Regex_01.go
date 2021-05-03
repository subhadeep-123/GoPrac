package main

import (
	"fmt"
	"regexp"
)

func findStringRe() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("facebook.com"))
}

func findstringIndex() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("abc.org"))
	fmt.Println(re.FindStringIndex("facebook.com"))
}

func findstringSubmatch() {
	re := regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re.FindStringSubmatch("flying"))
	fmt.Println(re.FindStringSubmatch("abcdfloatingxyz"))
	fmt.Println(re.FindStringSubmatch("asfdflewing.com"))

}
func main() {
	findStringRe()
	findstringIndex()
	findstringSubmatch()
}
