package main

import "fmt"

func main() {
	fmt.Println(CoinsCombination(10))  // [10]
	fmt.Println(CoinsCombination(11))  // [10 1]
	fmt.Println(CoinsCombination(125)) // [100 10 10 5]
}

func CoinsCombination(sum int) []int {
	var comb []int
	coins := []int{100, 50, 10, 5, 1}
	if sum < 0 {
		return comb // возвращаем пустой массив, если сумма отрицательная
	}
	for _, coin := range coins {
		for sum >= coin {
			comb = append(comb, coin)
			sum -= coin
		}
	}
	return comb
}
