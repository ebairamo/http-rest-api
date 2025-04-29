package main

import "fmt"

// Функция для проверки, является ли число простым
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	// Проверяем делимость на все числа от 2 до num-1
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Функция для нахождения n-го простого числа
func NthPrime(n int) int {
	count := 0 // Счётчик найденных простых чисел
	num := 2   // Начинаем с 2-го числа, т.к. 2 - первое простое число

	for {
		if isPrime(num) {
			count++
			if count == n {
				return num // Возвращаем n-е простое число
			}
		}
		num++
	}
}

func main() {
	fmt.Println(NthPrime(1))  // 2
	fmt.Println(NthPrime(5))  // 11
	fmt.Println(NthPrime(10)) // 29
}
