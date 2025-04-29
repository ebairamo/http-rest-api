package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%q\n", NthWord("salem alem!", 1))  // "salem"
	fmt.Printf("%q\n", NthWord("how are you?", 4)) // ""
}

func NthWord(text string, n int) string {
	var textSt string
	var textStr []string
	var result string

	for i := 0; i < len(text)+1; i++ {
		if i == len(text) {
			textStr = append(textStr, textSt)
			break
		}
		textSt += string(text[i])
		if text[i] == ' ' {
			textStr = append(textStr, textSt)
			textSt = ""
		}
	}
	if len(textStr) < n {
		result = ""
		return result
	}
	result = textStr[n-1]
	// if len(textStr) > n {
	// 	result = textStr[n]
	// } else {
	// 	result = ""
	// }

	return result
}
