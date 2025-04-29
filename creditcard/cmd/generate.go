package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// validateCardNumber выполняет валидацию номера карты.
// В данном случае функция всегда возвращает true, но сюда можно вставить реальную логику валидации.

func generateCard(cardNumber string) []string {
	var card []int
	var validCards []string
	var starCount int
	if isAllDigits(cardNumber) {
		fmt.Print("")
	} else {
		os.Exit(1)
	}
	fmt.Println("Генерация карт для шаблона:", cardNumber) // Отладочный вывод

	for _, char := range cardNumber {
		if char == '*' {
			starCount++
			card = append(card, -1) // Используем -1 как маркер для '*'

		} else {
			num, _ := strconv.Atoi(string(char))
			card = append(card, num)
		}
	}
	if starCount >= 5 {
		fmt.Println("Ошибка колличество звезд", starCount)
		os.Exit(1)
	}

	count := 0

	switch starCount {
	case 4:
		for a := 0; a <= 9; a++ {
			for b := 0; b <= 9; b++ {
				for c := 0; c <= 9; c++ {
					for d := 0; d <= 9; d++ {
						cardnew := ""
						starIndex := 0

						for _, ch := range card {
							if ch == -1 {
								switch starIndex {
								case 0:
									cardnew += strconv.Itoa(a)
								case 1:
									cardnew += strconv.Itoa(b)
								case 2:
									cardnew += strconv.Itoa(c)
								case 3:
									cardnew += strconv.Itoa(d)
								}
								starIndex++
							} else {
								cardnew += strconv.Itoa(ch)
							}
						}

						if validateCardNumber(cardnew) {
							// fmt.Printf("%s\n", cardnew)
							validCards = append(validCards, cardnew)
							count++
						}
					}
				}
			}
		}
	case 3:
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					cardnew := ""
					starIndex := 0

					for _, ch := range card {
						if ch == -1 {
							switch starIndex {
							case 0:
								cardnew += strconv.Itoa(b)
							case 1:
								cardnew += strconv.Itoa(c)
							case 2:
								cardnew += strconv.Itoa(d)
							}
							starIndex++
						} else {
							cardnew += strconv.Itoa(ch)
						}
					}

					if validateCardNumber(cardnew) {
						// fmt.Printf("%s\n", cardnew)
						validCards = append(validCards, cardnew)
						count++
					}
				}
			}
		}
	case 2:
		for c := 0; c <= 9; c++ {
			for d := 0; d <= 9; d++ {
				cardnew := ""
				starIndex := 0

				for _, ch := range card {
					if ch == -1 {
						switch starIndex {
						case 0:
							cardnew += strconv.Itoa(c)
						case 1:
							cardnew += strconv.Itoa(d)
						}
						starIndex++
					} else {
						cardnew += strconv.Itoa(ch)
					}
				}

				if validateCardNumber(cardnew) {
					// fmt.Printf("%s\n", cardnew)
					validCards = append(validCards, cardnew)
					count++
				}
			}
		}
	case 1:
		for d := 0; d <= 9; d++ {
			cardnew := ""

			for _, ch := range card {
				if ch == -1 {
					cardnew += strconv.Itoa(d)
				} else {
					cardnew += strconv.Itoa(ch)
				}
			}

			if validateCardNumber(cardnew) {
				// fmt.Printf("%s\n", cardnew)

				validCards = append(validCards, cardnew)
				count++
			}
		}
	}

	fmt.Println("Количество валидных карт:", count)
	return validCards
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && r != '*' {
			return false
		}
	}
	return true
}

// errors
// ./creditcard generate "4400430180漢字***"
