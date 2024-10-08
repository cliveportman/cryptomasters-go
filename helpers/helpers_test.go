package helpers

import (
	"testing"
)

func TestIsValidHex(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	if !isValidHex(hexString) {
		t.Error("Expected true, got false")
	}
}

func TestHexToBase64(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String, err := hexToBase64(hexString)
	if err != nil {
		t.Error(err)
	}
	if base64String != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("Expected SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t, got ", base64String)
	}
}

func TestSplitStringIntoLines(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyz"
	expected := "abcdefghij\nklmnopqrst\nuvwxyz"
	result := SplitStringIntoLines(text, 10)
	if result != expected {
		t.Error("Expected:\nabcdefghij\nklmnopqrst\nuvwxyz\n, got:\n", result)
	}
	
}

func TestXORCombinationWith2EqualLengthStrings(t *testing.T) {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	result, err := TwoStringsXOR(hexString1, hexString2)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Error("Expected 746865206b696420646f6e277420706c6179, got", result)
	}
}

func TestScoreText(t *testing.T) {
	text1 := "Some random sentence"
	text2 := "adlna!e;oHFHFHw;d;ws"
	score1 := scoreText(text1)
	score2 := scoreText(text2)
	if score1 <= score2 {
		t.Error("Expected score1 to be higher than score2")
	}
}

func TestCreateKeyForRepeatingKeyXOR(t *testing.T) {
	key := "QUACK"
	length := 10
	expected := "QUACKQUACK"
	result := CreateKeyForRepeatingKeyXOR(key, length)
	if result != expected {
		t.Error("Expected QUACKQUACK, got", result)
	}
}
