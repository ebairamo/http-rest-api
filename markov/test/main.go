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

func main() {
	var result []string
	var cmd string
	var prefix string
	var count int
	if len(os.Args) > 2 {
		cmd = os.Args[1] // начальная команда
		counts := os.Args[2]
		coun, err := strconv.Atoi(counts) // количество слов для генерации
		result = append(result, cmd)
		if err != nil {
			fmt.Println("Ошибка преобразования количества слов:", err)
			return
		}

		prefix = cmd
		// Создаем map для префиксов и суффиксов
		prefixMap := make(map[string][]string)

		// Открываем файл text.txt
		file, err := os.Open("text.txt")
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			return
		}
		defer file.Close()

		// Читаем файл построчно
		scanner := bufio.NewScanner(file)
		var prevWords []string // для хранения слов с предыдущей строки
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line) // Разбиваем строку на отдельные слова

			// Если есть слова с предыдущей строки, объединяем их с текущими
			if len(prevWords) > 0 {
				words = append(prevWords, words...)
			}

			// Проверяем, что есть достаточно слов для создания префиксов и суффиксов
			if len(words) >= 3 {
				// Итерируем по словам, формируем префиксы и суффиксы
				for i := 0; i < len(words)-2; i++ {
					prefix := words[i] + " " + words[i+1] // Префикс - два подряд идущих слова
					suffix := words[i+2]                  // Суффикс - слово, следующее за префиксом

					// Заполняем map для префиксов
					prefixMap[prefix] = append(prefixMap[prefix], suffix)
				}

				// Обновляем prevWords для случая разрыва строки
				prevWords = words[len(words)-2:]
			} else {
				// Если недостаточно слов, сохраняем их для следующей строки
				prevWords = append([]string{}, words...)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
		}

		for count < coun {
			// Проверяем, существует ли префикс
			if suffixes, found := prefixMap[prefix]; found {
				suffixes = removeDuplicates(suffixes)
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				next := suffixes[r.Intn(len(suffixes))]
				// Добавляем следующий суффикс в результат
				result = append(result, next)
				cmd = cmd + " " + next
				firstSpaceIndex := strings.Index(cmd, " ")

				// Оставляем только последние два слова в команде
				if firstSpaceIndex != -1 {
					cmd = cmd[firstSpaceIndex+1:]
					prefix = cmd
					count++
				}

			} else {
				fmt.Printf("Префикс '%s' не найден.\n", prefix)
				break // Прерываем цикл, если префикс не найден
			}
		}
	}
	fmt.Println(result)
}

func removeDuplicates(s []string) []string {
	unique := make(map[string]struct{})
	result := []string{}

	for _, v := range s {
		if _, exists := unique[v]; !exists {
			unique[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
