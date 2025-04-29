package main

import "fmt"

func main() {
	fmt.Println(Repeat("abc", 3)) // "abcabcabc"
	fmt.Println(Repeat("123", 2)) // "123123"
}

func Repeat(s string, count int) string {
	Repeat := ""
	for i := 0; i < count; i++ {
		Repeat += s
	}
	return Repeat
}
