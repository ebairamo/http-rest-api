package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateCard(generatedCardNumber string) []string {
	var card []int
	var validCards []string
	var starCount int

	for _, char := range generatedCardNumber {
		if char == '*' {
			starCount++
		} else {
			num, _ := strconv.Atoi(string(char))
			card = append(card, num)
		}
	}
	count := 0
	switch {
	case starCount == 4:
		for a := 0; a <= 9; a++ {
			for b := 0; b <= 9; b++ {
				for c := 0; c <= 9; c++ {
					for d := 0; d <= 9; d++ {
						// Создаем новую строку cardnew для каждой комбинации
						cardnew := ""

						// Добавляем цифры из исходной карты
						for _, ch := range card {
							if ch == -1 {
								// Если это место для звезды, пропустим
								cardnew += "0" // или другой подходящий символ, если это нужно
							} else {
								cardnew += strconv.Itoa(ch)
							}
						}

						// Добавляем новые цифры в строку
						cardnew += strconv.Itoa(a)
						cardnew += strconv.Itoa(b)
						cardnew += strconv.Itoa(c)
						cardnew += strconv.Itoa(d)

						// Валидация и вывод результата
						if validateCard(cardnew) {
							fmt.Printf("%s\n", cardnew)
							count++
						}
					}
				}
			}
		}

	case starCount == 3:

		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					// Создаем новую строку cardnew для каждой комбинации
					cardnew := ""

					// Добавляем цифры из исходной карты
					for _, ch := range card {
						if ch == -1 {
							// Если это место для звезды, пропустим
							cardnew += "0" // или другой подходящий символ, если это нужно
						} else {
							cardnew += strconv.Itoa(ch)
						}
					}

					// Добавляем новые цифры в строку

					cardnew += strconv.Itoa(b)
					cardnew += strconv.Itoa(c)
					cardnew += strconv.Itoa(d)

					// Валидация и вывод результата
					if validateCard(cardnew) {
						fmt.Printf("%s\n", cardnew)
						count++
					}
				}
			}
		}

	case starCount == 2:

		for c := 0; c <= 9; c++ {
			for d := 0; d <= 9; d++ {
				// Создаем новую строку cardnew для каждой комбинации
				cardnew := ""

				// Добавляем цифры из исходной карты
				for _, ch := range card {
					if ch == -1 {
						// Если это место для звезды, пропустим
						cardnew += "0" // или другой подходящий символ, если это нужно
					} else {
						cardnew += strconv.Itoa(ch)
					}
				}

				// Добавляем новые цифры в строку

				cardnew += strconv.Itoa(c)
				cardnew += strconv.Itoa(d)

				// Валидация и вывод результата
				if validateCard(cardnew) {
					fmt.Printf("%s\n", cardnew)
					count++
				}
			}
		}

	case starCount == 1:

		for d := 0; d <= 9; d++ {
			// Создаем новую строку cardnew для каждой комбинации
			cardnew := ""

			// Добавляем цифры из исходной карты
			for _, ch := range card {
				if ch == -1 {
					// Если это место для звезды, пропустим
					cardnew += "0" // или другой подходящий символ, если это нужно
				} else {
					cardnew += strconv.Itoa(ch)
				}
			}

			// Добавляем новые цифры в строку

			cardnew += strconv.Itoa(d)

			// Валидация и вывод результата
			if validateCard(cardnew) {
				fmt.Printf("%s\n", cardnew)
				count++
			}
		}

		// for a := 0; a < 10; a++ {
		// 	for b := 0; b < 10; b++ {
		// 		for c := 0; c < 10; c++ {
		// 			for d := 0; d < 10; d++ {
		// 				cardnew := append([]int(nil), card...)
		// 				cardnew = append(cardnew, a, b, c, d)

		// even = even[:0]
		// odd = odd[:0]
		// resultodd := 0
		// resulteven := 0

		// for i := 0; i < len(cardnew); i++ {0

		// 	if i%2 == 0 {
		// 		odd = append(odd, cardnew[i])
		// 	} else {
		// 		even = append(even, cardnew[i])
		// 	}
		// }

		// for i := 0; i < len(odd); i++ {
		// 	multiply := odd[i] * 2
		// 	if multiply > 9 {
		// 		multiply -= 9
		// 	}

		// 	resultodd += multiply
		// }

		// for i := 0; i < len(even); i++ {
		// 	resulteven += even[i]
		// }

		// resultfinal := resultodd + resulteven
		// if resultfinal%10 == 0 {
		// 	cardStr := fmt.Sprint(cardnew)
		// 	ValidCards = append(ValidCards, strings.Join(strings.Fields(cardStr), ""))
		// }

		// cardnew = nil
	}
	fmt.Println(count)
	return validCards
}

func validateCard(cardNumber string) bool {
	var odd []int
	var even []int
	var card []int
	var resultodd int
	var resulteven int
	var resultfinal int

	for _, char := range strings.TrimSpace(cardNumber) {
		str := string(char)
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка преобразования символа:", err)
			return false
		}
		card = append(card, num)
	}

	if len(card) < 13 {
		fmt.Println("Длина карты меньше 13 цифр")
		return false

	}

	for i := 0; i < len(card); i++ {
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

	return resultfinal%10 == 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'generate', 'validate', 'information', or 'issue' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		if len(os.Args) < 3 {
			fmt.Println("expected an argument for 'generate'")
			os.Exit(1)
		}
		cardNumber := os.Args[2]
		generateCard(cardNumber)
		handleGenerate(cardNumber)

	case "validate":
		if len(os.Args) < 3 {
			fmt.Println("expected an argument for 'validate'")
			os.Exit(1)
		}
		cardNumber := os.Args[2]

		validateCard(cardNumber)
		handleValidate(cardNumber)
		if validateCard(cardNumber) {
			fmt.Println("OK")
		} else {
			fmt.Println("Incorrect")
		}

	case "information":
		if len(os.Args) < 3 {
			fmt.Println("expected an argument for 'information'")
			os.Exit(1)
		}
		cardNumber := os.Args[2]
		handleInformation(cardNumber)

	case "issue":
		if len(os.Args) < 3 {
			fmt.Println("expected an argument for 'issue'")
			os.Exit(1)
		}
		cardNumber := os.Args[2]
		handleIssue(cardNumber)

	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}

// Временные заглушки для функций
func handleGenerate(cardNumber string) {
	// Временная заглушка: ничего не делает
}

func handleValidate(cardNumber string) {
	// Временная заглушка: ничего не делает
}

func handleInformation(cardNumber string) {
	// Временная заглушка: ничего не делает
}

func handleIssue(cardNumber string) {
	// Временная заглушка: ничего не делает
}
