package challenges

import (
	"cryptomasters/helpers"
	"fmt"
	"os"
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

	content = []byte(strings.ReplaceAll(string(content), "\n", ""))
	


	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	// text1 := "this is a test"
	// text2 := "wokka wokka!!!"
	// result, err := helpers.HammingDifference(text1, text2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	// fmt.Printf("Hamming difference between %s and %s is %d\n", text1, text2, result)

	// We're trying to find the keysize and using possible values of 2 to 40
	for i := 2; i < 41; i++ {
		// We're taking the first two blocks of size i
		block1 := content[:i]
		block2 := content[i:i*2]
		// bytes1, err := base64.StdEncoding.DecodeString(block1)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// bytes2, err := base64.StdEncoding.DecodeString(block2)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// We're calculating the hamming distance between the two blocks
		result, err := helpers.HammingDifference(block1, block2)
		if err != nil {
			fmt.Println(err)
		}
		// return the average hamming distance as a float
		average := float64(result)/float64(i)
		fmt.Printf("Hamming difference is %f\n",average)
	}
}
