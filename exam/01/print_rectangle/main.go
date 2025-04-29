package main

import "github.com/alem-platform/ap"

func main() {
	PrintRectangle(2, 3)
	// Output:
	// 0 0 0
	// 0 0 0
}

func PrintRectangle(a, b int) {
	for j := 0; j < a; j++ {
		for i := 0; i < b; i++ {
			ap.PutRune('0')
			if i != b-1 {
				ap.PutRune(' ')
			}
		}
		ap.PutRune('?')
		ap.PutRune('\n')
	}
}
