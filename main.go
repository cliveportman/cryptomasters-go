package main

import (
	"cryptomasters/helpers"
	"fmt"
	"strconv"
)

func main() {
	h := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	b, err := helpers.HexToBytes(h)
	if err != nil {
		fmt.Println(err)
		return
	}

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	res := make([]string, len(chars))
	for ri, r := range chars {
		x := make([]byte, len(b))
		for i := range b {
			x[i] = b[i] ^ byte(r) // c is of type rune, so we need to cast it to byte
		}
		res[ri] = string(x)
		fmt.Println(strconv.Itoa(ri) + "(" + string(r) + ", " + strconv.Itoa(helpers.ScoreText(string(x))) + "):" + string(x))
	}
	fmt.Println("Result incoming...")
	fmt.Println(res[49])
}
