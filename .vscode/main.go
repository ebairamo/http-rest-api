package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result []string
	var count int

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		counts := os.Args[2]
		coun, err := strconv.Atoi(counts)
		fmt.Println(coun)
		// Открываем файл
		file, err := os.Open("text.txt")
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			return
		}
		defer file.Close()

		// Читаем файл построчно
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Разбиваем строку на слова
			words := strings.Fields(scanner.Text())

			// Создаем map для хранения двух слов как ключ и третьего слова как значение
			mark := make(map[string]string)

			// Проходим по словам и берем первые два слова как ключ, третье как значение
			for i := 0; i < len(words)-2; i++ {
				key := words[i] + " " + words[i+1] // Создаем ключ из двух слов
				mark[key] = words[i+2]             // Третье слово — это значение
			}
			if count < coun {
				// Проверяем, существует ли команда в качестве ключа
				if nextWord, found := mark[cmd]; found {
					fmt.Println(nextWord, mark[cmd])
					result = append(result, nextWord)
					fmt.Println(result)
					count++
					nextWord = ""
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
		}
	}
}
