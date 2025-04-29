package main

import (
	"fmt"
	"io"
	"os"
)

type Pixel struct {
	B byte
	G byte
	R byte
}

// Функция для поворота изображения на 90, 180 или 270 градусов
func Rotate(pixelData *[]Pixel, width int, height int, angle int) (newWidth, newHeight int) {
	newPixelData := make([]Pixel, width*height)

	switch angle {
	case 90:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				newX := height - 1 - y
				newY := x
				newPixelData[newY*height+newX] = (*pixelData)[y*width+x]
			}
		}
		newWidth, newHeight = height, width

	case 180:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				newX := width - 1 - x
				newY := height - 1 - y
				newPixelData[newY*width+newX] = (*pixelData)[y*width+x]
			}
		}
		newWidth, newHeight = width, height

	case 270:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				newX := y
				newY := width - 1 - x
				newPixelData[newY*height+newX] = (*pixelData)[y*width+x]
			}
		}
		newWidth, newHeight = height, width
	}

	*pixelData = newPixelData
	return newWidth, newHeight
}

func main() {
	file, err := os.Open("sample.bmp")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	header := make([]byte, 54)
	_, err = io.ReadFull(file, header)
	if err != nil {
		fmt.Println("Ошибка при чтении заголовка:", err)
		return
	}

	width := int(header[18]) | int(header[19])<<8 | int(header[20])<<16 | int(header[21])<<24
	height := int(header[22]) | int(header[23])<<8 | int(header[24])<<16 | int(header[25])<<24
	fmt.Printf("Ширина: %d, Высота: %d\n", width, height)

	// Рассчитываем размер строки с учетом выравнивания
	rowSize := (width*3 + 3) &^ 3
	rawPixelData := make([]byte, rowSize*height)
	_, err = io.ReadFull(file, rawPixelData)
	if err != nil {
		fmt.Println("Ошибка при чтении данных изображения:", err)
		return
	}

	// Преобразуем сырые байты в пиксели
	pixelData := make([]Pixel, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*rowSize + x*3
			pixelData[y*width+x] = Pixel{
				B: rawPixelData[pos],
				G: rawPixelData[pos+1],
				R: rawPixelData[pos+2],
			}
		}
	}

	// Вызов поворота изображения на 90 градусов
	newWidth, newHeight := Rotate(&pixelData, width, height, 270)

	// Открываем файл для записи
	outFile, err := os.Create("output.bmp")
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outFile.Close()

	// Обновляем заголовок с новыми размерами
	header[18] = byte(newWidth)
	header[19] = byte(newWidth >> 8)
	header[20] = byte(newWidth >> 16)
	header[21] = byte(newWidth >> 24)
	header[22] = byte(newHeight)
	header[23] = byte(newHeight >> 8)
	header[24] = byte(newHeight >> 16)
	header[25] = byte(newHeight >> 24)

	_, err = outFile.Write(header)
	if err != nil {
		fmt.Println("Ошибка при записи заголовка:", err)
		return
	}

	// Рассчитываем новый размер строки с выравниванием
	newRowSize := (newWidth*3 + 3) &^ 3
	padding := make([]byte, newRowSize-newWidth*3)

	// Записываем данные пикселей обратно в файл, добавляя байты выравнивания
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			pixel := pixelData[y*newWidth+x]
			outFile.Write([]byte{pixel.B, pixel.G, pixel.R})
		}
		// Добавляем байты выравнивания после каждой строки
		outFile.Write(padding)
	}

	fmt.Println("Изображение успешно сохранено!")
}
