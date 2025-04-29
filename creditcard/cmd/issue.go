package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleIssue(brands string, issuers string, brand string, issuer string) {
	// fmt.Println("Brands file:", brands)
	// fmt.Println("Issuers file:", issuers)
	// fmt.Println("Selected Brand:", brand)
	// fmt.Println("Selected Issuer:", issuer)

	var fields []string

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

		brandSerch := strings.TrimSpace(fields[0])

		if brandSerch == brand {
			// fmt.Println(fields[1])
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Could not read file:", err)
		os.Exit(1)
	}
	// }
	// // Название карты
	// for i, v := range sixNumbers {
	// 	strSlice[i] = strconv.Itoa(v)

	// 	sixDigitsStr := strings.Join(strSlice, "")

	file, err = os.Open(issuers)

	// Создаем сканер для построчного чтения файла
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fieldse := strings.Split(line, ":")

		if len(fieldse) < 2 {
			fmt.Println("Неверный формат строки:", line)
			continue
		}

		issuerSearch := strings.TrimSpace(fieldse[0])
		if issuerSearch == issuer {
			issuerCard := fieldse[1] // Инициализируем строку
			// fmt.Println("lloo", fieldse[1])
			// fmt.Println(issuerSearch)

			// Генерация случайных чисел
			for i := 0; i < 6; i++ {
				randomNumber := generateRandomNumber()
				str := strconv.Itoa(randomNumber)
				issuerCard = issuerCard + str

				// fmt.Println(randomNumber)
				// fmt.Println(issuerCard)
			}
			result := issuerCard // Объединяем строки
			starsapp := "****"
			result += starsapp
			// Выводим содержимое среза
			// fmt.Println(starsapp)

			// fmt.Println(result)

			lol := generateCard(result)

			fmt.Println(lol[666])

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Scanner encountered an error:", err)
		os.Exit(1)
	}
}

func generateRandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(10) // Возвращаем случайное число от 0 до 9
}
