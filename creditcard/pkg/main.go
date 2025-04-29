package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cardNumber := "4400430180300003"
	var card []int
	var sixNumbers []int
	var issuer string
	var brand string

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
	fmt.Println(card)
	for i := 0; i < 6; i++ {
		sixNumbers = append(sixNumbers, card[i])
	}
	fmt.Println(sixNumbers)
	// Преобразуем каждый элемент в строку
	for i, v := range sixNumbers {
		strSlice[i] = strconv.Itoa(v)

		sixDigitsStr := strings.Join(strSlice, "")

		// Открываем файл
		file, err := os.Open("brands.txt")
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

		file, err := os.Open("issuers.txt")
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
	fmt.Println(brand)
	fmt.Println(issuer)
}
