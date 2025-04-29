package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleInformation(brands string, issuers string, brand string, issuer string) {
	cardNumbers := os.Args[4:]

	for _, cardNumber := range cardNumbers {
		fmt.Println(cardNumber)

		// Проверка валидности номера карты
		if validateCardNumber(cardNumber) {
			fmt.Println("Correct: yes")
		} else {
			fmt.Println("Correct: no")
			fmt.Println("Card Brand: -")
			fmt.Println("Card Issuer: -")
			return
		}
		var card []int
		var sixNumbers []int

		strSlice := make([]string, 6)
		var fields []string
		for _, char := range strings.TrimSpace(cardNumber) {

			num, _ := strconv.Atoi(string(char))
			// if err != nil {
			// 	fmt.Println("Ошибка преобразования символа:", err)
			// 	fmt.Println(str)
			// 	return false
			// }
			card = append(card, num)
		}

		for i := 0; i < 6; i++ {
			sixNumbers = append(sixNumbers, card[i])
		}

		// Преобразуем каждый элемент в строку
		for i, v := range sixNumbers {
			strSlice[i] = strconv.Itoa(v)

			sixDigitsStr := strings.Join(strSlice, "")

			// Открываем файл
			file, err := os.Open(brands)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Could not open file:", err)
				os.Exit(1)
			}
			defer file.Close()

			// Создаем сканер для построчного чтения файла
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				// Считываем строку
				line := scanner.Text()
				fields = strings.Split(line, ":")

				sixNum := strings.TrimSpace(fields[1])
				// fmt.Println(fields[1])
				if sixNum == sixDigitsStr {
					brand = fields[0]
				}

			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Could not read file:", err)
				os.Exit(1)
			}
		}
		// Название карты
		for i, v := range sixNumbers {
			strSlice[i] = strconv.Itoa(v)

			sixDigitsStr := strings.Join(strSlice, "")

			file, err := os.Open(issuers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Could not open file:", err)
				os.Exit(1)
			}
			defer file.Close()

			// Создаем сканер для построчного чтения файла
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				// Считываем строку
				line := scanner.Text()
				fieldse := strings.Split(line, ":")

				sixNum := strings.TrimSpace(fieldse[1])
				// fmt.Println(fields[1])
				if sixNum == sixDigitsStr {
					issuer = fieldse[0]
				}

			}
		}
		fmt.Println("Card Brand:", brand)
		fmt.Println("Card Issuer:", issuer)
	}
}
