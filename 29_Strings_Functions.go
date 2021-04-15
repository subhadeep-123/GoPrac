package main

import (
	"fmt"
	"strings"
)

func iterateStr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
}

func printASCII(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println("s"[i])
	}
}

func toUpper() {
	str := "india"
	fmt.Println(strings.ToUpper(str))
}

func printString(str string) {
	fmt.Println(len(str))
}

func toLower() {
	str := "india"
	fmt.Println(strings.ToLower(str))
}

func hasPrefix() {
	str := "india"
	fmt.Println(strings.HasPrefix(str, "in"))
}

func hasSuffix() {
	str := "india"
	fmt.Println(strings.HasSuffix(str, "ia"))
}

func Join() {
	var arr = []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings.Join(arr, "*"))
}

func Repeat() {
	str := "Repeat"
	fmt.Println(strings.Repeat(str, 2))
}

func checkContains() {
	str := "Hi...there"
	fmt.Println(strings.Contains(str, "th"))
}

func checkIndex() {
	str := "Hi...there"
	fmt.Println(strings.Index(str, "th"))
}

func checkCount() {
	str := "Hi...there"
	fmt.Println(strings.Count(str, "e"))
}

func doReplace() {
	str := "eaeaeaeaea"
	fmt.Println(strings.Replace(str, "e", "Z", 2))
	fmt.Println(str)
	fmt.Println(strings.ReplaceAll(str, "e", "L"))
}

func doSplit() {
	str := "Hey How you doing man"
	var arr []string = strings.Split(str, " ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s\t", arr[i])
	}
	fmt.Println("%q\n", strings.Split(str, " "))
}

func doTrim() {
	fmt.Println(strings.TrimSpace("\t\n I love my \ncountry \n\t\r\n"))
}

func checkContainsAny() {
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))
}
func doCompare() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Compare("f", "a"))
}

func main() {
	str := "I love my country"
	printString(str)
	iterateStr(str)
	// printASCII(str)
	toUpper()
	toLower()
	hasPrefix()
	hasSuffix()
	Join()
	Repeat()
	checkContains()
	checkIndex()
	checkCount()
	doReplace()
	doSplit()
	doCompare()
	doTrim()
	checkContainsAny()
}
