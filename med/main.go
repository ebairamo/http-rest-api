package main

import "fmt"

func main() {
	fmt.Println(ArrayMin([]int{10, -1, 20, 4, 5})) // -1
	fmt.Println(ArrayMin([]int{15, 7, 8, 12}))     // 7
	fmt.Println(ArrayMin([]int{}))                 // 0
}
func ArrayMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := nums[0]

	for i := 0; i < len(nums); i++ {

		// fmt.Println(nums[i], nums[j])
		if nums[i] < count {
			count = nums[i]
		}

	}

	return count
}
