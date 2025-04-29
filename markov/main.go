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
	var count int
	var savedLine string // Переменная для сохранения строки

	// Проверяем, есть ли аргументы командной строки
	if len(os.Args) > 2 {
		cmd := os.Args[1] // начальная команда
		counts := os.Args[2]
		coun, err := strconv.Atoi(counts) // количество слов для генерации
		result = append(result, cmd)
		if err != nil {
			fmt.Println("Ошибка преобразования количества слов:", err)
			return
		}

		// Открываем файл и читаем его, пока count < coun
		preResult := []string{}
		for count < coun {
			file, err := os.Open("text.txt")
			if err != nil {
				fmt.Println("Ошибка при открытии файла:", err)
				return
			}

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				// Читаем первую строку
				line1 := scanner.Text()

				// Проверяем, есть ли еще строки, если нет, выходим из цикла
				if !scanner.Scan() {
					break
				}

				// Читаем вторую строку
				line2 := scanner.Text()

				// Если вторая строка пустая, сохраняем первую строку
				if line2 == "" {
					savedLine = line1
					continue // Переходим к следующей итерации
				}

				// Если savedLine не пустая, объединяем её с line2
				if savedLine != "" {
					line1 = savedLine
					savedLine = "" // Очищаем сохраненную строку
				}

				// Объединяем две строки
				line := line1 + " " + line2
				words := strings.Fields(line) // Разбиваем строку на слова

				if strings.Contains(strings.ToLower(line), strings.ToLower(cmd)) {
					for i := 0; i < len(words)-2; i++ {
						word := words[i] + " " + words[i+1]
						if strings.ToLower(word) == strings.ToLower(cmd) {
							if i+2 < len(words) {
								nextWord := words[i+2]

								// Добавляем nextWord в preResult, если его там еще нет
								if !contains(preResult, nextWord) {
									preResult = append(preResult, nextWord)
								}
							}
						}
					}
				}
			}

			file.Close() // Закрываем файл после чтения

			// Генерируем результат
			if len(preResult) > 0 {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				next := preResult[r.Intn(len(preResult))]
				preResult = []string{}
				if !contains(result, next) {
					result = append(result, next)
				}
				// Обновляем команду, добавляя новое слово
				cmd = cmd + " " + next
				firstSpaceIndex := strings.Index(cmd, " ")

				// Оставляем только последние два слова в команде
				if firstSpaceIndex != -1 {
					cmd = cmd[firstSpaceIndex+1:]
					cmd = removeNonSpaceChars(cmd)
				}
				count++ // Увеличиваем счетчик
			} else {
				break
			}

		}
		fmt.Println(result)
	}
}

// Функция для проверки, содержится ли слово в срезе
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func removeNonSpaceChars(s string) string {
	// Удаление всех символов, кроме букв, цифр и пробелов
	return strings.ReplaceAll(s, "\"", "")
}
