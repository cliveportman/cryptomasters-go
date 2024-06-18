package challenges

import (
	"cryptomasters/helpers"
	"encoding/hex"
	"fmt"
)

/*
Provided with a string of text, convert it to bytes, then encrypt it using a repeating key XOR.
*/
func Challenge5() {

	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	textHex := hex.EncodeToString([]byte(text))

	key := "ICE"
	repeatingKey := helpers.CreateKeyForRepeatingKeyXOR(key, len(text))
	repeatingKeyHex := hex.EncodeToString([]byte(repeatingKey))
	result, error := helpers.TwoStringsXOR(textHex, repeatingKeyHex)
	if error != nil {
		fmt.Println(error)
		return
	}
	result = helpers.SplitStringIntoLines(result, 75)
	fmt.Println(result)
}
