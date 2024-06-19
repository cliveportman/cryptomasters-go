package challenges

import (
	"cryptomasters/helpers"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
A text file containing a number of strings is given. One of the strings has been encrypted using single-character XOR. Find it.
*/
func Challenge4() {
	content, error := os.ReadFile("assets/c4.txt")
	if error != nil {
		fmt.Println(error)
	}
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	results := make([]helpers.Result, len(lines))
	for i, line := range lines {
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

	fmt.Printf("Char: %s, Score: %d, Text: %s\n", results[0].Character, results[0].Score, results[0].Text)
}
