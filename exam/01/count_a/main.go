package main

import "fmt"

func main() {
	fmt.Println(CountA([]string{"apple", "banana", "avocado"})) // 6
	fmt.Println(CountA([]string{"xyz", "uvw", "rst"}))          // 0
	fmt.Println(CountA([]string{"aardvark", "alpaca", "antA"})) // 7
}

func CountA(slice []string) int {
	count := 0
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 'a' || slice[i][j] == 'A' {
				count++
			}
		}
	}
	return count
}
