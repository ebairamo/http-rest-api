package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	// Определяем команды и флаги
	if len(os.Args) < 2 {
		fmt.Println("Использование: <команда> [флаги] <входной файл> <выходной файл>")
		fmt.Println("Доступные команды: apply, header")
		return
	}

	command := os.Args[1]

	// Флаги для команды `apply`
	mirrorFlag := flag.String("mirror", "", "Отразить изображение: horizontal или vertical")
	var filters stringSlice
	flag.Var(&filters, "filter", "Применить фильтр (можно несколько): red, green, blue, negative, pixelate")
	var rotateFlags stringSlice
	flag.Var(&rotateFlags, "rotate", "Поворот: right (90° вправо), left (90° влево)")
	cropFlag := flag.String("crop", "", "Crop parameters in the format: offsetX-offsetY or offsetX-offsetY-width-height")

	flag.CommandLine.Parse(os.Args[2:])

	// Проверка команды
	if command != "apply" {
		fmt.Println("Неизвестная команда:", command)
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Использование: <команда> [флаги] <входной файл> <выходной файл>")
		return
	}

	inputFile := args[0]
	var outputFile string

	if len(args) < 2 {
		outputFile = "output.bmp"
	} else {
		outputFile = args[1]
	}

	// Чтение BMP файла
	width, height, header, pixelData, err := readBMP(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении BMP:", err)
		return
	}

	var bmpData BMPdata // Убедитесь, что BMPdata правильно определена

	fmt.Printf("Исходное изображение: ширина=%d, высота=%d\n", width, height) // Исправлено: используем width и height

	if *cropFlag != "" {
		cropParams := strings.Split(*cropFlag, ",")
		err := CropImage(&bmpData, cropParams) // Здесь bmpData должна содержать данные изображения
		if err != nil {
			fmt.Println("Ошибка обрезки:", err)
			return
		}
		fmt.Println("Изображение обрезано успешно!")
	} else {
		fmt.Println("Не указаны параметры обрезки.")
	}

	// Применение зеркалирования
	if *mirrorFlag != "" {
		mirrorImage(pixelData, width, height, *mirrorFlag)
	}

	// Применение фильтров
	for _, filter := range filters {
		switch filter {
		case "red":
			applyRedFilter(pixelData)
		case "green":
			applyGreenFilter(pixelData, width, height, 0, 0)
		case "blue":
			applyBlueFilter(pixelData, width, height, 0, 0)
		case "negative":
			applyNegativeFilter(pixelData, width, height, 0, 0)
		case "pixelate":
			applyPixelateFilter(pixelData, width, height, 10)
		case "blur":
			applyBlurFilter(pixelData, width, height, 10)
		}
	}

	for _, rotate := range rotateFlags {
		switch rotate {
		case "90", "right":
			newWidth, newHeight := Rotate(&pixelData, width, height, 90)
			width, height = newWidth, newHeight // Обновляем текущие ширину и высоту
			updateHeader(header, width, height) // Обновляем заголовок
		case "180":
			width, height = Rotate(&pixelData, width, height, 180)
		case "270", "left":
			newWidth, newHeight := Rotate(&pixelData, width, height, 270)
			width, height = newWidth, newHeight // Обновляем текущие ширину и высоту
			updateHeader(header, width, height) // Обновляем заголовок
		}
	}

	// Запись в выходной файл
	err = writeBMP(outputFile, header, pixelData, width, height)
	if err != nil {
		fmt.Println("Ошибка при записи BMP:", err)
		return
	}
	fmt.Println("Изображение успешно обработано и сохранено в", outputFile)
}

// Пример функций для фильтров, поворотов и зеркалирования

// Функции для чтения и записи BMP
func updateHeader(header []byte, width int, height int) {
	header[18] = byte(width)
	header[19] = byte(width >> 8)
	header[20] = byte(width >> 16)
	header[21] = byte(width >> 24)
	header[22] = byte(height)
	header[23] = byte(height >> 8)
	header[24] = byte(height >> 16)
	header[25] = byte(height >> 24)
}
