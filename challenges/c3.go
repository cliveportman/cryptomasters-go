package challenges

import (
	"cryptomasters/helpers"
	"fmt"
)

/*
A hex-encoded string has been XOR'd against a single character. Find the key, decrypt the message.
*/
func Challenge3() {
	inputHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	result, error := helpers.SingleCharacterXOR(inputHex)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("Char: %s, Score: %d, Text: %s\n", result.Character, result.Score, result.Text)
}
