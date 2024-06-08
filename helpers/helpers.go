package helpers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func IsValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// HexToBase64 Converts a hex string to a base64 string
func HexToBase64(h string) (string, error) {
	// If no hex string has been provided, return an error
	if !IsValidHex(h) {
		return "", fmt.Errorf("%s is not a valid hex string", h)
	}
	// convert to bytes first
	bytes, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	// then convert bytes to base64
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String, nil
}

// HexToBytes Converts a hex string then to a byte slice
func HexToBytes(h string) ([]byte, error) {
	if !IsValidHex(h) {
		return nil, fmt.Errorf("%s is not a valid hex string", h)
	}
	b, err := hex.DecodeString(h)
	if err != nil {
		return nil, fmt.Errorf("could not decode hex string: %v", err)
	}
	return b, nil
}

// XORCompareTwoHexStrings XORs two hex strings (of equal length) together, returning a hex string
func XORCompareTwoHexStrings(buffer1 string, buffer2 string) (string, error) {
	// The buffers have to be of equal length
	if len(buffer1) != len(buffer2) {
		return "", fmt.Errorf("buffers are not of equal length")
	}
	bytes1, err := HexToBytes(buffer1)
	if err != nil {
		return "", err
	}
	bytes2, err := HexToBytes(buffer2)
	if err != nil {
		return "", err
	}
	// Create a slice of the correct length to store the result
	resultBytes := make([]byte, len(bytes1))
	// Then perform the XOR operation on each byte, adding the result to the slice
	for i := range bytes1 {
		resultBytes[i] = bytes1[i] ^ bytes2[i]
	}
	return hex.EncodeToString(resultBytes), nil
}

// Note the use of a space at the beginning of the string - in Challenge 3, this made a big difference
var englishLetterFrequency = []rune(" etaoinshrdlcumwfgypbvkjxqz") // Most common characters first
var challenge3HintFrequency = []rune("ETAOIN SHRDLU") // The hint text... which returns the same result. Feels like cheating.

// ScoreText scores a string based on the frequency of English characters
func ScoreText(text string) int {
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
