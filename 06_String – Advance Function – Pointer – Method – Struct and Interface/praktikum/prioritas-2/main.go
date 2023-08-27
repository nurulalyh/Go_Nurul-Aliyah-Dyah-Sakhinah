package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var newStr string = ""
	// var lowercase string = strings.ToLower(input)

	for _, val := range input {
		if val >= 'a' && val <= 'z' {
			char := 'a' + (val-'a'+rune(offset))%26
			newStr += string(char)
		} else if val >= 'A' && val <= 'Z' {
			char := 'A' + (val-'A'+rune(offset))%26
			newStr += string(char)
		} else {
			newStr += string(val)
		}
	}
	return newStr
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
