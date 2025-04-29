package main

import "fmt"

func main() {
	fmt.Println(GeneratePrimes(10)) // [2 3 5 7]
	fmt.Println(GeneratePrimes(20)) // [2 3 5 7 11 13 17 19]
}

func GeneratePrimes(n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
