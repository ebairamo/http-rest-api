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
	count := 1
	var result []string
	preResult := []string{}
	var finalResult []string

	if len(os.Args) > 1 {

		cmd := os.Args[1]
		result = append(result, cmd)
		counts := os.Args[2]
		coun, err := strconv.Atoi(counts)
		fmt.Println(coun)
		if err != nil {
			// ... handle error
			panic(err)
		}
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
			// Выводим каждое слово с его индексом + 1

			mark := make(map[string]string)

			// Заполняем мапу последовательностями слов
			for i := 0; i < len(words)-1; i++ {
				mark[words[i]] = words[i+1]
			}
			if count < coun {

				for key, value := range mark {
					// Проверяем, совпадает ли значение с cmd
					if key == cmd {
						// Добавляем ключ в preResult, если значение равно cmd
						fmt.Println("key", key)
						fmt.Println(cmd)
						// fmt.Println(value)
						preResult = append(preResult, cmd, value)
						fmt.Println("pre", preResult)

					}
				}

				// cmd = resultOne

				// Выводим итоговый результат

				// rand.Seed(time.Now().UnixNano())

				if len(preResult) > 0 {
					rand.Seed(time.Now().UnixNano())
					resultOne := preResult[rand.Intn(len(preResult)-1)]
					// fmt.Println("ты", resultOne)
					// Случайный индекс от 0 до len(preResult)-1
					for _, resulted := range result {
						// fmt.Println(resulted, "sdfsdfdfsgfgsdf")
						if resultOne != resulted {
							// fmt.Println("dsfsdfsfd", resultOne, resulted)
							if !contains(result, resultOne) {
								result = append(result, cmd)
								result = append(result, resultOne)
								finalResult = append(finalResult, resultOne)
								cmd = resultOne
							}
						}

					}
					fmt.Println(count)
					count++
					preResult = nil
				}

			}
		}
		fmt.Println("len", finalResult)
	}

}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
