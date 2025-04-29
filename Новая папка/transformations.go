package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DIBHeader struct {
	// other fields
	Width     int32
	Height    int32
	ImageSize uint32 // Add this field to store the image size in bytes
	// other fields
}

// Структура BMP-данных
type BMPdata struct {
	BMP   BMPHeader
	DIB   DIBHeader
	Color [][]Pixel
}

// Функция зеркального отображения
func mirrorImage(pixelData []Pixel, width, height int, direction string) {
	switch direction {
	case "horizontal", "h", "horizontally", "hor":
		for y := 0; y < height; y++ {
			for x := 0; x < width/2; x++ {
				oppX := width - 1 - x
				pixelData[y*width+x], pixelData[y*width+oppX] = pixelData[y*width+oppX], pixelData[y*width+x]
			}
		}
	case "vertical", "v", "vertically", "ver", "venera":
		for y := 0; y < height/2; y++ {
			for x := 0; x < width; x++ {
				originalIndex := y*width + x
				mirroredIndex := (height-1-y)*width + x
				pixelData[originalIndex], pixelData[mirroredIndex] = pixelData[mirroredIndex], pixelData[originalIndex]
			}
		}
	default:
		fmt.Println("Неверное направление для зеркалирования. Используйте: horizontal, vertical.")
	}
}

// Функция поворота изображения
func Rotate(pixelData *[]Pixel, width int, height int, angle int) (int, int) {
	var newWidth, newHeight int
	if angle == 180 {
		newWidth, newHeight = width, height
	} else {
		newWidth, newHeight = height, width
	}
	newPixelData := make([]Pixel, newWidth*newHeight)

	switch angle {
	case 90:
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				newX := i
				newY := width - 1 - j
				newPixelData[newY*newWidth+newX] = (*pixelData)[i*width+j]
			}
		}
		width, height = newWidth, newHeight

	case 180:
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				newX := width - 1 - j
				newY := height - 1 - i
				newPixelData[newY*newWidth+newX] = (*pixelData)[i*width+j]
			}
		}

	case 270:
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				newX := height - 1 - i
				newY := j
				newPixelData[newY*newWidth+newX] = (*pixelData)[i*width+j]
			}
		}
		width, height = newWidth, newHeight

	default:
		fmt.Println("Unsupported rotation angle. Use 90, 180, or 270 degrees.")
		return width, height
	}

	*pixelData = newPixelData
	return width, height
}

// Функция для обрезки изображения
func CropImage(bmpData *BMPdata, crops []string) error {
	for _, cropParams := range crops {
		if err := applyCrop(bmpData, cropParams); err != nil {
			return err
		}
	}
	return nil
}

// Функция для обработки параметров обрезки
func applyCrop(bmpData *BMPdata, cropParams string) error {
	params := strings.Split(cropParams, "-")
	if len(params) != 2 && len(params) != 4 {
		return fmt.Errorf("недопустимые параметры обрезки: %s", cropParams)
	}

	// Чтение начальных координат
	offsetX, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("недопустимый сдвиг X: %v", err)
	}

	offsetY, err := strconv.Atoi(params[1])
	if err != nil {
		return fmt.Errorf("недопустимый сдвиг Y: %v", err)
	}

	// Определение ширины и высоты в зависимости от количества параметров
	var width, height int
	if len(params) == 4 {
		width, err = strconv.Atoi(params[2])
		if err != nil {
			return fmt.Errorf("недопустимая ширина: %v", err)
		}

		height, err = strconv.Atoi(params[3])
		if err != nil {
			return fmt.Errorf("недопустимая высота: %v", err)
		}
	} else {
		// Если не указаны размеры, используем оставшиеся размеры изображения
		width = int(bmpData.DIB.Width) - offsetX
		height = int(bmpData.DIB.Height) - offsetY
	}

	// Проверка допустимых значений координат и размеров
	if offsetX < 0 || offsetY < 0 || width <= 0 || height <= 0 ||
		offsetX+width > int(bmpData.DIB.Width) || offsetY+height > int(bmpData.DIB.Height) {
		return fmt.Errorf("размеры обрезки превышают размеры изображения")
	}

	// Логика обрезки: создание новой области пикселей
	newColor := make([][]Pixel, height)
	for i := 0; i < height; i++ {
		newColor[i] = make([]Pixel, width)
		copy(newColor[i], bmpData.Color[offsetY+i][offsetX:offsetX+width])
	}

	// Обновление данных BMP с учетом новых размеров
	bmpData.Color = newColor
	bmpData.DIB.Width = int32(width)
	bmpData.DIB.Height = int32(height)
	bmpData.DIB.ImageSize = uint32(height * ((width*3 + 3) &^ 3)) // Выравнивание строки на 4 байта
	bmpData.BMP.FileSize = bmpData.BMP.DataOffset + bmpData.DIB.ImageSize

	// Информация о новых размерах после обрезки
	fmt.Printf("Новые размеры изображения: ширина=%d, высота=%d\n", bmpData.DIB.Width, bmpData.DIB.Height)

	return nil
}
