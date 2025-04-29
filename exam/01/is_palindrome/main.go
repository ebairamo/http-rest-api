package main

import "fmt"

func main() {
	fmt.Println(IsPalindrome("madam"))
	// Output: true

	fmt.Println(IsPalindrome("hello"))
	// Output: false
}

func IsPalindrome(s string) bool {
	var pali string
	for i := len(s) - 1; i >= 0; i-- {
		pali += string(s[i])
	}
	if s == pali {
		return true
	} else {
		return false
	}

}
  