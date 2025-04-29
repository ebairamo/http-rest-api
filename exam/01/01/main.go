package main

import "fmt"

func main() {
	fmt.Println(Trim("   Salem student!   "))
	fmt.Println(Trim("  a   a  a"))
}

func Trim(s string) string {
	var trim string
	for i := 0; i < len(s)-1; i++ {

		if s[i] != ' ' {
			trim = trim + string(s[i])
		}
		if s[i] == ' ' && s[i+1] != ' ' {

			trim = trim + string(s[i])

		}

	}

	return trim

}

// func Trim(s string) string {
// 	var trim string
// 	for i := 0; i < len(s)-1; i++ {
// 		if s[i] == ' ' {

// 			if s[i+1] != ' ' {
// 				trim = trim + string(s[i])
// 			}
// 		}
// 	}
// 	return trim
// }
