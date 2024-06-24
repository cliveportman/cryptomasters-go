package challenges

import (
	"cryptomasters/helpers"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
Provided with a string of text, convert it to bytes, then encrypt it using a repeating key XOR.
*/
func Challenge6() {
	content, error := os.ReadFile("assets/c4.txt")
	if error != nil {
		fmt.Println(error)
	}
	// Looks like the content is one single base64 string with line breaks (only one example of padding, at the end)
	// so just remove the line breaks and convert to bytes
	content = []byte(strings.ReplaceAll(string(content), "\n", ""))
	results := []helpers.HammingDistanceResult{}
	// Trying key sizes from 2 to 40 - the key size producing the lowest Hamming distance is likely to be the actual key size
	for keysize := 2; keysize < 41; keysize++ {
		// Take the first two blocks of keysize length
		block1 := content[:keysize]
		block2 := content[keysize : keysize*2]
		// And calculate the hamming difference
		result, err := helpers.HammingDifference(block1, block2)
		if err != nil {
			fmt.Println(err)
		}
		// Normalise them by dividing by the keysize (converting to float for greater accuracy)
		average := float64(result) / float64(keysize)
		results = append(results, helpers.HammingDistanceResult{KeySize: keysize, Score: average})
	}
	// Sort the results so the lowest Hamming difference is top and print that
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score < results[j].Score
	})
	fmt.Printf("Keysize: %d, Hamming difference: %f\n", results[0].KeySize, results[0].Score)
	keySize := results[0].KeySize

	// Break the content into blocks of keySize
	var blocks [][]byte
	for i := 0; i < len(content); i += keySize {
		end := i + keySize
		if end > len(content) {
			end = len(content)
		}
		blocks = append(blocks, content[i:end])
	}
	// Transpose the blocks. We'll end up with a slice of length keysize, each item of which is a slice of length len(blocks)
	transposedBlocks := make([][]byte, keySize)
	// Loop through the keysize
	for i := 0; i < keySize; i++ {
		// Loop through the blocks
		for _, block := range blocks {
			// If the keysize is less than the length of the block
			if i < len(block) { // Check if the block is long enough
				transposedBlocks[i] = append(transposedBlocks[i], block[i])
			}
		}
	}
	// XOR through each block, returning the highest scoring result
	XORresults := make([]helpers.Result, len(transposedBlocks))
	for i, block := range transposedBlocks {
		result, error := helpers.SingleCharacterXORBytes(block)
		if error != nil {
			fmt.Println(error)
		} else {
			XORresults[i] = result
		}

	}
	// Create an empty slice to hold the key, then populate it with the highest scoring character from each block
	keyBase64 := []string{}
	for _, XORResult := range XORresults {
		keyBase64 = append(keyBase64, XORResult.Character)
		//fmt.Printf("Char: %s, Score: %d, Text: %s\n", XORResult.Character, XORResult.Score, XORResult.Text)
	}
	// Converting the key to bytes
	keyBytes := [][]byte{}
	for i, key := range keyBase64 {
		keyBytes = append(keyBytes, []byte(key))
		fmt.Printf("%d: %b or %s\n", i, keyBytes[i], keyBytes[i])
	}
	// for i, key := range keyBytes {
	// 	fmt.Printf("%d: %b or %s\n", i, key, key)
	// }

}
