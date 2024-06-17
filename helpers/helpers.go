package helpers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func isValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// HexToBytes Converts a hex string to a slice of bytes
func hexToBytes(h string) ([]byte, error) {
	if !isValidHex(h) {
		return nil, fmt.Errorf("%s is not a valid hex string", h)
	}
	b, err := hex.DecodeString(h)
	if err != nil {
		return nil, fmt.Errorf("could not decode hex string: %v", err)
	}
	return b, nil
}

// HexToBase64 Converts a hex string to a base64 string
func hexToBase64(h string) (string, error) {
	// convert to bytes
	bytes, err := hexToBytes(h)
	if err != nil {
		return "", err
	}
	// ...then convert bytes to base64
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String, nil
}

// StringToBytes Converts a string to a slice of bytes
func StringToBytes(s string) []byte {
	return []byte(s)
}

// ScoreText scores a string based on the frequency of English characters
func scoreText(text string) int {
	score := 0
	for _, char := range strings.ToLower(text) {
		for i, freqChar := range englishLetterFrequency {
			if char == freqChar {
				// We subtract the index from the length so that more frequent letters have a higher score
				score += len(englishLetterFrequency) - i
				break
			}
		}
	}
	return score
}

// Create key from repeating string
func CreateKeyForRepeatingKeyXOR(key string, length int) string {
	var result string
	for len(result) < length {
		result += key
	}
	return result[:length]
}

// func CreateKeyForRepeatingKeyXOR(key string, length int) []byte {
// 	keyBytes := StringToBytes(key)
// 	keyLength := len(keyBytes)
// 	result := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		result[i] = keyBytes[i%keyLength]
// 	}
// 	return result
// }

// TwoStringsXOR XORs two hex strings (of equal length) together, returning a hex string
func TwoStringsXOR(buffer1 string, buffer2 string) (string, error) {
	if len(buffer1) != len(buffer2) {
		return "", fmt.Errorf("buffers are not of equal length")
	}
	bytes1, err := hexToBytes(buffer1)
	if err != nil {
		return "", err
	}
	bytes2, err := hexToBytes(buffer2)
	if err != nil {
		return "", err
	}
	// Create a slice of the correct length to store the results, then iterate over each byte with the XOR operation
	resultBytes := make([]byte, len(bytes1))
	for i := range bytes1 {
		resultBytes[i] = bytes1[i] ^ bytes2[i]
	}
	return hex.EncodeToString(resultBytes), nil
}

// SingleCharacterXOR XORs a string against a list of single characters, returning the highest scoring result
func SingleCharacterXOR(inputHex string) (Result, error) {
	inputBytes, err := hexToBytes(inputHex)
	if err != nil {
		fmt.Println(err)
		return Result{}, err
	}
	results := make([]Result, len(base64Characters))

	for charIndex, charRune := range base64Characters {
		x := make([]byte, len(inputBytes))
		for i := range inputBytes {
			x[i] = inputBytes[i] ^ byte(charRune) // charRune is of type rune, so we need to cast it to byte
		}
		results[charIndex] = Result{
			Character: string(charRune),
			Score:     scoreText(string(x)),
			Text:      string(x),
		}
	}
	// Sort the results so the highest scoring characters' results are top
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	// Return the top scoring result
	return results[0], nil

}
