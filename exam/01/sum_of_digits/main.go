package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumOfDigits(123))  // 6
	fmt.Println(SumOfDigits(-456)) // 15
	fmt.Println(SumOfDigits(0))    // 0
}

func SumOfDigits(n int) int {
	zero := 0
	// Если число равно нулю, сразу возвращаем строку "0"
	if n == 0 {
		return zero
	}

	var result string // Переменная для хранения результата
	sign := 1         // Переменная для сохранения знака числа (по умолчанию положительное)

	// Если число отрицательное, делаем его положительным и помечаем знак
	if n < 0 {
		sign = -1
		n *= -1 // Делаем число положительным
	}

	// Извлекаем цифры из числа
	for n > 0 {
		digit := n % 10                           // Получаем последнюю цифру числа
		result = string('0'+rune(digit)) + result // Преобразуем цифру в символ и добавляем в начало строки
		n /= 10                                   // Убираем последнюю цифру из числа
	}

	// Если число было отрицательным, добавляем минус в начало строки
	if sign == -1 {
		result = "-" + result
	}
	var final int
	var minus string
	for i := 0; i < len(result); i++ {
		if int(result[i]) == '-' {
			minus = string(result[i])
			minus = minus
		} else {
			final += (int(result[i]) - '0')
		}
	}
	return final

}
