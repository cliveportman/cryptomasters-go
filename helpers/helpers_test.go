package helpers

import (
	"testing"
)

func TestIsValidHex(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	if !IsValidHex(hexString) {
		t.Error("Expected true, got false")
	}
}

func TestHexToBase64(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String, err := HexToBase64(hexString)
	if err != nil {
		t.Error(err)
	}
	if base64String != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("Expected SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t, got ", base64String)
	}
}

func TestXORCombination(t *testing.T) {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	result, err := XORCompareTwoHexStrings(hexString1, hexString2)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Error("Expected 746865206b696420646f6e277420706c6179, got", result)
	}
}
