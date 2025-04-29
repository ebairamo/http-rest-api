// main.go

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Открытие CSV файла
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Используем наш CSVParser
	var csvparser CSVParser = &superCSVParser{}

	// Чтение строк из файла
	for {
		line, err := csvparser.ReadLine(file)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}

		fmt.Println("Line:", line)
		fmt.Println("Number of fields:", csvparser.GetNumberOfFields())

		// Выводим все поля
		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			field, err := csvparser.GetField(i)
			if err != nil {
				fmt.Println("Error getting field:", err)
				continue
			}
			fmt.Printf("Field %d: %s\n", i, field)
		}

	}
}
