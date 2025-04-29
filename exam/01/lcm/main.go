package main

import "fmt"

func main() {
	fmt.Println(LCM(4, 6))  // 12
	fmt.Println(LCM(21, 6)) // 42
	fmt.Println(LCM(5, 7))  // 35
}

func LCM(a, b int) int {
	var ap []int
	var bp []int
	var result int
	out := false
	for i := 1; i < 20; i++ {
		ap = append(ap, a*i)
		bp = append(bp, b*i)
	}
	for i := 0; i < len(ap); i++ {
		if out == false {
			for j := 0; j < len(bp); j++ {
				if ap[i] == bp[j] {
					result = ap[i]
					out = true
				}
			}
		}
	}
	return result
}
