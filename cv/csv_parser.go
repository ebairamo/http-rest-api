package main

import (
	"errors"
	"io"
)

// Определение ошибок
var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

// Интерфейс CSVParser
type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

// Структура для хранения состояния парсера
type superCSVParser struct {
	lastLine string   // Последняя прочитанная строка
	fields   []string // Поля последней строки
}

// Реализация метода ReadLine с использованием io.Reader
func (p *superCSVParser) ReadLine(r io.Reader) (string, error) {
	var line string
	var inQuotes bool
	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(line) > 0 {
					break
				}
				return "", io.EOF
			}
			return "", err
		}

		// Добавляем символ в строку
		line += string(buf)

		// Проверка на кавычки
		if buf[0] == '"' {
			inQuotes = !inQuotes
		}

		// Прекращаем чтение строки, если не находимся внутри кавычек и встречаем конец строки
		if (buf[0] == '\n' || buf[0] == '\r') && !inQuotes {
			if buf[0] == '\r' {
				// Чтение для \r\n в Windows
				_, _ = r.Read(buf)
			}
			break
		}
	}

	// Проверка на некорректные кавычки
	if countQuotes(line)%2 != 0 {
		return "", ErrQuote
	}

	p.lastLine = trim(line) // Убираем пробелы вручную
	p.fields = parseCSVLine(p.lastLine)
	return p.lastLine, nil
}

// trim удаляет пробелы в начале и конце строки
func trim(s string) string {
	// Убираем пробелы в начале
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\t') {
		start++
	}

	// Убираем пробелы в конце
	end := len(s) - 1
	for end >= start && (s[end] == ' ' || s[end] == '\t') {
		end--
	}

	return s[start : end+1]
}

// Подсчет количества кавычек
func countQuotes(s string) int {
	count := 0
	for _, char := range s {
		if char == '"' {
			count++
		}
	}
	return count
}

// Парсинг CSV строки вручную
func parseCSVLine(line string) []string {
	var fields []string
	var field string
	inQuotes := false

	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch == '"' {
			if inQuotes && i+1 < len(line) && line[i+1] == '"' {
				// Экранированные кавычки (двойные кавычки внутри кавычек)
				field += `"`
				i++ // Пропустить следующую кавычку
			} else {
				inQuotes = !inQuotes
			}
		} else if ch == ',' && !inQuotes {
			fields = append(fields, field)
			field = ""
		} else {
			field += string(ch)
		}
	}
	fields = append(fields, field) // добавляем последнее поле
	return fields
}

// Реализация метода GetField
func (p *superCSVParser) GetField(n int) (string, error) {
	if n < 0 || n >= len(p.fields) {
		return "", ErrFieldCount
	}
	return p.fields[n], nil
}

// Реализация метода GetNumberOfFields
func (p *superCSVParser) GetNumberOfFields() int {
	return len(p.fields)
}
