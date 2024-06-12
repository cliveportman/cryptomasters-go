package challenges

import (
	"cryptomasters/helpers"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

/*
A hex-encoded string has been XOR'd against a single character. Find the key, decrypt the message.
 */
func Challenge4() {
	content, error := ioutil.ReadFile("assets/strings.txt")
	if error != nil {
		fmt.Println(error)
	}
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	fmt.Println(strconv.Itoa(len(lines)))

	results := make([]helpers.Result, len(lines))
	for i, line := range lines {
		//fmt.Println(strconv.Itoa(i) + ":" +line)
		result, error := helpers.SingleCharacterXOR(line)
		if error != nil {
			fmt.Println(error)
		} else {
			results[i] = result
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	//for _, res := range results {
	//	fmt.Printf("Char: %c, Score: %d, Text: %s\n", res.Character, res.Score, res.Text)
	//}
	fmt.Printf("Char: %c, Score: %d, Text: %s\n", results[0].Character, results[0].Score, results[0].Text)
}
