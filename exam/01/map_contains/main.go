package main

import "fmt"

func main() {
	ascending := func(a, b int) bool {
		return a < b
	}

	descending := func(a, b int) bool {
		return a > b
	}

	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	Sort(arr, ascending)
	fmt.Println(arr) // [1 1 2 3 4 5 5 6 9]

	arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 16}
	Sort(arr, descending)
	fmt.Println(arr) // [9 6 5 5 4 3 2 1 1]
}
func Sort(arr []int, fn func(int, int) bool) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if fn(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
