package main

import (
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
	var allword []string
	var showHelp bool
	prefixMap := make(map[string][]string)

	maxWords := 100
	prefixLength := 2
	startPrefix := ""
	showHelp = false

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--help":
			showHelp = true
		case "-w":
			if i+1 < len(os.Args) {
				counts, err := strconv.Atoi(os.Args[i+1])
				if err != nil || counts <= 0 || counts > 10000 {
					fmt.Println("Ошибка: Неверное количество слов.")
					return
				}
				maxWords = counts
				i++
			} else {
				fmt.Println("Ошибка: Отсутствует значение для -w.")
				return
			}
		case "-p":
			if i+1 < len(os.Args) {
				startPrefix = os.Args[i+1]
				cmd = startPrefix
				i++
			} else {
				fmt.Println("Ошибка: Отсутствует значение для -p.")
				return
			}
		case "-l":
			if i+1 < len(os.Args) {
				length, err := strconv.Atoi(os.Args[i+1])
				if err != nil || length < 1 || length > 5 {
					fmt.Println("Ошибка: Неверная длина префикса.")
					return
				}
				if prefixLength != 0 {
					prefixLength = length
				}
				i++
			} else {
				fmt.Println("Ошибка: Отсутствует значение для -l.")
				return
			}
		}
	}
	if showHelp {
		printHelp()
		return
	}

	prefix = cmd
	// Указываем путь к файлу

	// Открываем файл
	file := os.Stdin

	// Получаем размер файла
	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("Ошибка при получении информации о файле: %v\n", err)
		return
	}
	result = append(result, cmd)
	// Создаем буфер для чтения
	content := make([]byte, stat.Size())

	// Читаем файл целиком
	_, err = file.Read(content)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}

	// Преобразуем содержимое в строку
	strContents := string(content)

	// Разбиваем строку на слова
	words := strings.Fields(strContents)

	// Выводим все слова
	allword = append(allword, words...)

	if len(allword) >= prefixLength+1 {
		// Итерируем по словам, формируем префиксы и суффиксы
		for i := 0; i <= len(allword)-prefixLength-1; i++ {
			prefixSlice := allword[i : i+prefixLength] // Префикс - несколько подряд идущих слов
			prefix := strings.Join(prefixSlice, " ")   // Соединяем слова в строку
			suffix := allword[i+prefixLength]          // Суффикс - слово, следующее за префиксом

			// Заполняем map для префиксов
			prefixMap[prefix] = append(prefixMap[prefix], suffix)
		}
		if maxWords <= 0 {
			fmt.Println("Ошибка: Неверное количество слов.")
			return
		}

		maxWords = maxWords - prefixLength
		// Обновляем prevWords для случая разрыва строки
		if startPrefix == "" {
			for i := 0; i < maxWords+1; i++ {
				result = append(result, allword[i])
			}

		} else {
			for count < maxWords {
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
				}
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

func printHelp() {
	fmt.Println("Markov Chain text generator.")
	fmt.Println("Usage: markovchain [-w <N>] [-p <S>] [-l <N>]")
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words.")
	fmt.Println("  -p S    Starting prefix.")
	fmt.Println("  -l N    Prefix length.")
}
