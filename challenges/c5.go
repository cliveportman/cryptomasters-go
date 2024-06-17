package challenges

import (
	"cryptomasters/helpers"
	"encoding/hex"
	"fmt"
	"strings"
)

/*
Provided with a string of text, convert it to bytes, then encrypt it using a repeating key XOR.
*/
func Challenge5() {
	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	lines := strings.Split(string(text), "\n")

	textHex := hex.EncodeToString([]byte(text))
	key := "ICE"
	repeatingKey := helpers.CreateKeyForRepeatingKeyXOR(key, len(text))
	repeatingKeyHex := hex.EncodeToString([]byte(repeatingKey))
	fmt.Println(len(text))
	fmt.Println(len(repeatingKey))
	result, error := helpers.TwoStringsXOR(textHex, repeatingKeyHex)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(result)

	// textBytes := helpers.StringToBytes(text)
	// XORkey := helpers.CreateKeyForRepeatingKeyXOR(key, len(textBytes))

	// // encrypt the text using the XORkey
	// encryptedText := helpers.RepeatingKeyXOR(textBytes, XORkey)

	// fmt.Println(textBytes)
	// fmt.Println(XORkey)
}
