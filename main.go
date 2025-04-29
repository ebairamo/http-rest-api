package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var odd []int
	var even []int
	var card []int
	var resultodd int
	var resulteven int
	var resultfinal int

	for _, cardNumber := range os.Args[1:] {

		for _, char := range cardNumber {

			str := string(char)

			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Ошибка преобразования символа:", err)
				return
			}

			card = append(card, num)
		}
		if len(card) < 13 {
			fmt.Println("Длинна карты меньше 13 цифр")
		}
		fmt.Println(card)
		for i := 0; i < (len(card)); i++ {
			if i%2 == 0 {
				odd = append(odd, card[i])
			} else {
				even = append(even, card[i])
			}
		}

		for i := 0; i < len(odd); i++ {
			multiply := odd[i] * 2
			if multiply > 9 {
				multiply -= 9
			}

			resultodd += multiply

		}
		for i := 0; i < len(even); i++ {
			resulteven += even[i]
		}
		resultfinal = resultodd + resulteven
	}

	if resultfinal%10 == 0 {
		fmt.Println("Ok")
		fmt.Println(resultfinal)
		os.Exit(1)

	} else {
		fmt.Println("Incorrect")
		fmt.Println("final", resultfinal)
		os.Exit(0)
	}
}
