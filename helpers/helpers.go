package helpers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func IsValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

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

func XORCombination(buffer1 string, buffer2 string) (string, error) {
	// The buffers have to be of equal length
	if len(buffer1) != len(buffer2) {
		return "", fmt.Errorf("Buffers are not of equal length")
	}
	// Convert to base64 first
	buffer1Base64, err := HexToBase64(buffer1)
	if err != nil {
		return "", err
	}
	buffer2Base64, err := HexToBase64(buffer2)
	if err != nil {
		return "", err
	}
	// Then to bytes
	buffer1Bytes, err := base64.StdEncoding.DecodeString(buffer1Base64)
	if err != nil {
		return "", fmt.Errorf("invalid base64 string: %s", buffer1Base64)
	}
	buffer2Bytes, err := base64.StdEncoding.DecodeString(buffer2Base64)
	if err != nil {
		return "", fmt.Errorf("invalid base64 string: %s", buffer2Base64)
	}
	// Create a slice of the correct length to store the result
	resultBytes := make([]byte, len(buffer1Bytes))
	// Then perform the XOR operation on each byte, adding the result to the slice
	for i := range buffer1Bytes {
		resultBytes[i] = buffer1Bytes[i] ^ buffer2Bytes[i]
	}
	// Return the hex encoded result
	return hex.EncodeToString(resultBytes), nil
}
