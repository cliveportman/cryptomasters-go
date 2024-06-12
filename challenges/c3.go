package challenges

import (
	"cryptomasters/helpers"
	"fmt"
)
//type result struct {
//	character string
//	score int
//	text string
//}

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
	fmt.Printf("###########\nChar: %s\nScore: %d\nText: %s\n", result.Character, result.Score, result.Text)
	//for _, res := range results {
	//	fmt.Printf("Char: %c, Score: %d, Text: %s\n", res.character, res.score, res.text)
	//}
}
