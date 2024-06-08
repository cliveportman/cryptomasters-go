package challenges

import (
	"cryptomasters/helpers"
	"fmt"
	"sort"
)
type result struct {
	character string
	score int
	text string
}

/*
A hex-encoded string has been XOR'd against a single character. Find the key, decrypt the message.
 */
func Challenge3() {
	inputHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputBytes, err := helpers.HexToBytes(inputHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Interesting: the lowercase letters return the same as the uppercase letters but with invalid characters instead of spaces
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	results := make([]result, len(chars))

	for charIndex, charRune := range chars {
		x := make([]byte, len(inputBytes))
		for i := range inputBytes {
			x[i] = inputBytes[i] ^ byte(charRune) // charRune is of type rune, so we need to cast it to byte
		}
		results[charIndex] = result{
			character: string(charRune),
			score:     helpers.ScoreText(string(x)),
			text:      string(x),
		}
	}

	// Sort the results so the highest scoring is first
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	fmt.Println("Result incoming...\n##################")
	fmt.Printf("Char: %s\nScore: %d\nText: %s\n", results[0].character, results[0].score, results[0].text)
	//for _, res := range results {
	//	fmt.Printf("Char: %c, Score: %d, Text: %s\n", res.character, res.score, res.text)
	//}
}
