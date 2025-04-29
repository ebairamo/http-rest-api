package main

import "github.com/alem-platform/ap"

func main() {
	PrintAlphaPyramid(36)
	// Output:
	// A
	// B C
	// D E F
	// G H I J
	// K L M N O
	// P Q R S T U
	// V W X Y Z A B
	// C D E F G H I J
}
func PrintAlphaPyramid(n int) {
	var j int
	m := 1
	d := 1
	for i := 0; i < n; i++ {
		j = i
		if i >= 26 {
			j = i - 26
		}
		r := 'A' + rune(j)

		ap.PutRune(r)
		if d != m {
			ap.PutRune(' ')
		}
		if d == m {
			ap.PutRune('\n')
			m++
			d = 0
		}
		d++
	}
}
