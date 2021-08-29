package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "JavaTpoint@12345!@#$%^&*()"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Value to be encode " + data)
	fmt.Println("Encoded Value " + sEnc)
	fmt.Println("Value to be decoded " + sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Decoded Value " + string(sDec))
	url := "https://golang.org/ref/spec"
	fmt.Println("url to be encode  " + url)
	urlEncode := b64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded url  " + urlEncode)

	fmt.Println("value to be decode  " + urlEncode)
	strDecode2, _ := b64.URLEncoding.DecodeString(urlEncode)

	fmt.Println("Decoded value  " + string(strDecode2))
}
