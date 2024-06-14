package helpers

var base64Characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Note the use of a space at the beginning of the string - in Challenge 3, this made a big difference
var englishLetterFrequency = []rune(" etaoinshrdlcumwfgypbvkjxqz") // Most common characters first
var challenge3HintFrequency = []rune("ETAOIN SHRDLU")              // The hint text... which returns the same result. Feels like cheating.
