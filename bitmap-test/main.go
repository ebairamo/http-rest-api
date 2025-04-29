package main

import (
	"fmt"
	"io"
	"os"
)

// Структура для представления пикселя
type Pixel struct {
	B byte // Синий
	G byte // Зелёный
	R byte // Красный
}

// Функция для установки цвета пикселя по координатам (x, y)
func setPixel(pixelData []Pixel, width, x, y int, r, g, b byte) {
	pos := y*width + x
	pixelData[pos] = Pixel{B: b, G: g, R: r}
}

func mirrorRow(pixelData []Pixel, width int) {
	for y := 0; y < len(pixelData)/width; y++ {
		start := y * width
		end := start + width - 1
		for x := 0; x < width/2; x++ {
			pixelData[start+x], pixelData[end-x] = pixelData[end-x], pixelData[start+x]
		}
	}
}

func blue(pixelData []Pixel, width, height, x, y int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentB := pixelData[y*width+x].B

			pixelData[y*width+x] = Pixel{
				B: currentB,
				G: 0,
				R: 0,
			}
		}
	}
}

func green(pixelData []Pixel, width, height, x, y int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentG := pixelData[y*width+x].G

			pixelData[y*width+x] = Pixel{
				B: 0,
				G: currentG,
				R: 0,
			}
		}
	}
}

func red(pixelData []Pixel, width, height, x, y int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentR := pixelData[y*width+x].R

			pixelData[y*width+x] = Pixel{
				B: 0,
				G: 0,
				R: currentR,
			}
		}
	}
}

func negative(pixelData []Pixel, width, height, x, y int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentR := pixelData[y*width+x].R
			currentG := pixelData[y*width+x].G
			currentB := pixelData[y*width+x].B

			pixelData[y*width+x] = Pixel{
				B: 255 - currentB,
				G: 255 - currentG,
				R: 255 - currentR,
			}
		}
	}
}

func pix(pixelData []Pixel, width, height int) []Pixel {
	scale := 10000

	// Новый массив для увеличенных пикселей
	newWidth := width * scale
	newHeight := height * scale
	newPixelData := make([]Pixel, newWidth*newHeight)

	// Обработка каждого пикселя
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Получаем цвет текущего пикселя
			currentColor := pixelData[y*width+x]

			// Заполнение области размером scale x scale
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					bigX := x*scale + dx
					bigY := y*scale + dy

					// Устанавливаем цвет в новом массиве
					newPixelData[bigY*newWidth+bigX] = currentColor
				}
			}
		}
	}

	return newPixelData
}

// Теперь newPixelData содержит увеличенные пиксели

// Теперь newPixelData содержит увеличенные пиксели

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

	pixelData := make([]Pixel, width*height)
	rowSize := (width*3 + 3) &^ 3
	rawPixelData := make([]byte, rowSize*height)
	_, err = io.ReadFull(file, rawPixelData)
	if err != nil {
		fmt.Println("Ошибка при чтении данных изображения:", err)
		return
	}

	// Конвертируем байтовый массив в массив пикселей
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
	// green(pixelData, width, height, 0, 0)
	// blue(pixelData, width, height, 0, 0)
	// red(pixelData, width, height, 0, 0)
	// mirrorRow(pixelData, width)
	// negative(pixelData, width, height, 0, 0)
	pix(pixelData, width, height)

	outFile, err := os.Create("output.bmp")
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outFile.Close()

	_, err = outFile.Write(header)
	if err != nil {
		fmt.Println("Ошибка при записи заголовка:", err)
		return
	}

	// Записываем изменённые пиксельные данные
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			_, err := outFile.Write([]byte{pixelData[y*width+x].B, pixelData[y*width+x].G, pixelData[y*width+x].R})
			if err != nil {
				fmt.Println("Ошибка при записи данных изображения:", err)
				return
			}
		}
	}

	fmt.Println("Изменённые пиксели успешно записаны в output.bmp")
}
