package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	pickFlag  = flag.Bool("pick", false, "Randomly pick a single entry")
	stdinFlag = flag.Bool("stdin", false, "Read input from standard input")
)

func generateCard(generatedCardNumber string) []string {
	var card []int
	var starCount int
	var validCards []string

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
		fmt.Println(cardNumber)
		if err != nil {
			fmt.Println("Ошибка преобразования символа:", err)
			fmt.Println(str)
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

	if resultfinal%10 == 0 {

	}
	return true
}

func main() {
	flag.Parse() // Обрабатываем флаги командной строки

	if len(flag.Args()) < 1 {
		fmt.Println("ожидалось 'generate', 'validate', 'information' или 'issue' как подкоманда")
		os.Exit(1)
	}

	var cardNumber string

	// Обработка ввода через stdin
	if *stdinFlag {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения из stdin:", err)
			os.Exit(1)
		}
		cardNumber = strings.TrimSpace(input)
	} else if len(flag.Args()) > 1 {
		cardNumber = flag.Args()[1]
	} else {
		fmt.Println("ожидался аргумент с номером карты или флаг --stdin")
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "generate":
		if len(flag.Args()) < 2 {
			fmt.Println("expected an argument for 'generate'")
			os.Exit(1)
		}
		cardNumber := flag.Args()[1]
		validCards := generateCard(cardNumber)
		for _, card := range validCards {
			fmt.Println(card)
		}

	case "validate":
		if validateCard(cardNumber) {
			fmt.Println("OK")
		} else {
			fmt.Println("Некорректно")
		}

	case "information":

	case "issue":

	default:
		fmt.Println("неизвестная команда")
		os.Exit(1)
	}
}
