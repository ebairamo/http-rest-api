package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%q\n", WordNumber("salem alem! salem", "salem")) // 1
	fmt.Printf("%q\n", WordNumber("salem alem!", "alem!"))       // 0
	fmt.Printf("%q\n", WordNumber("how are you?", "salem"))      // 0
}

func WordNumber(text string, word string) int {
	t := 0
	textFull := ""
	var textStr []string
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' {
			textFull = textFull + string(text[i])
		}
		if text[i] == ' ' || i == len(text)-1 {
			textStr = append(textStr, textFull)
			textFull = ""
		}

	}
	for j := 0; j < len(textStr); j++ {
		if textStr[j] == word {
			t++
		}
	}
	return t
}
