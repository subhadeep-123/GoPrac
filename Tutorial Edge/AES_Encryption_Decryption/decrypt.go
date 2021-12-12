package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

const FILENAME string = "encryptedData.txt"

func ReadFromFile(fname string) []byte {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data
}

func main() {
	key := []byte("passphrasewhichneesstobe32bytes!")
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	chiperData := ReadFromFile(FILENAME)

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	if len(chiperData) < nonceSize {
		fmt.Println(err)
	}
	nonce, chiperData := chiperData[:nonceSize], chiperData[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, chiperData, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(plainText))
}
