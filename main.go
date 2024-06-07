package main

import (
	"cryptomasters/helpers"
	"fmt"
)

func main() {
	hexString := "1c0111001f010100061a024b53535009181c"
	base64String, err := helpers.HexToBase64(hexString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(base64String)
}
