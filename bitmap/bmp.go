package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// readBMP читает BMP-файл и возвращает ширину, высоту, заголовок и данные пикселей
func readBMP(filename string) (int, int, []byte, []Pixel, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Чтение заголовка BMP (54 байта)
	header := make([]byte, 54)
	if _, err := io.ReadFull(file, header); err != nil {
		return 0, 0, nil, nil, fmt.Errorf("ошибка при чтении заголовка: %v", err)
	}

	// Проверка на корректный BMP файл
	if header[0] != 'B' || header[1] != 'M' {
		return 0, 0, nil, nil, fmt.Errorf("это не BMP файл")
	}

	// Извлекаем ширину и высоту из заголовка
	width := int(binary.LittleEndian.Uint32(header[18:22]))
	height := int(binary.LittleEndian.Uint32(header[22:26]))

	// Отладочный вывод
	fmt.Printf("Заголовок загружен. Ширина: %d, Высота: %d\n", width, height)

	// Чтение пикселей
	pixelData, err := readPixels(file, width, height)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("ошибка при чтении пикселей: %v", err)
	}

	return width, height, header, pixelData, nil
}

// writeBMP записывает данные BMP-файла, включая заголовок и пиксели
func writeBMP(filename string, header []byte, pixelData []Pixel, width, height int) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка при создании выходного файла: %v", err)
	}
	defer outFile.Close()

	// Запись заголовка
	if _, err := outFile.Write(header); err != nil {
		return fmt.Errorf("ошибка при записи заголовка: %v", err)
	}

	// Запись данных пикселей с учётом выравнивания строк
	rowSize := (width*3 + 3) &^ 3
	padding := make([]byte, rowSize-width*3)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := pixelData[y*width+x]
			if _, err := outFile.Write([]byte{pixel.B, pixel.G, pixel.R}); err != nil {
				return fmt.Errorf("ошибка при записи данных изображения: %v", err)
			}
		}
		// Добавляем отступы (padding)
		if _, err := outFile.Write(padding); err != nil {
			return fmt.Errorf("ошибка при записи отступов: %v", err)
		}
	}

	return nil
}

// readPixels читает пиксели BMP-файла и возвращает их как массив Pixel
func readPixels(file *os.File, width, height int) ([]Pixel, error) {
	rowSize := (width*3 + 3) &^ 3
	rawPixelData := make([]byte, rowSize*height)

	// Читаем сырые данные пикселей
	if _, err := io.ReadFull(file, rawPixelData); err != nil {
		return nil, fmt.Errorf("ошибка при чтении данных изображения: %v", err)
	}

	// Инициализация массива пикселей
	pixelData := make([]Pixel, width*height)

	// Преобразуем сырые байты в массив пикселей
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
	return pixelData, nil
}

// func ReadBMPHeader(file *os.File, bmpData *BMPdata) error {
// 	// Чтение BMP заголовка
// 	if err := binary.Read(file, binary.LittleEndian, &bmpData.BMP); err != nil {
// 		return fmt.Errorf("ошибка чтения BMP-заголовка: %v", err)
// 	}
// 	// Чтение DIB заголовка
// 	if err := binary.Read(file, binary.LittleEndian, &bmpData.DIB); err != nil {
// 		return fmt.Errorf("ошибка чтения DIB-заголовка: %v", err)
// 	}

// 	// Проверка, что ширина и высота корректны
// 	if bmpData.DIB.Width <= 0 || bmpData.DIB.Height <= 0 {
// 		return fmt.Errorf("некорректные размеры изображения: ширина=%d, высота=%d", bmpData.DIB.Width, bmpData.DIB.Height)
// 	}

// 	fmt.Printf("Ширина изображения: %d, Высота изображения: %d\n", bmpData.DIB.Width, bmpData.DIB.Height)
// 	return nil
// }

func ReadBMPHeaders(file *os.File, bmpData *BMPdata) error {
	// Reading BMP Header (first 14 bytes)
	var bmpHeader [14]byte
	_, err := file.Read(bmpHeader[:])
	if err != nil {
		return fmt.Errorf("error reading BMP header: %v", err)
	}
	bmpData.BMP.FileSize = int32(bmpHeader[2]) | int32(bmpHeader[3])<<8 | int32(bmpHeader[4])<<16 | int32(bmpHeader[5])<<24
	bmpData.BMP.DataOffset = int32(bmpHeader[10]) | int32(bmpHeader[11])<<8 | int32(bmpHeader[12])<<16 | int32(bmpHeader[13])<<24

	// Reading DIB Header (next part, usually 40 bytes for BITMAPINFOHEADER)
	var dibHeader [40]byte
	_, err = file.Read(dibHeader[:])
	if err != nil {
		return fmt.Errorf("error reading DIB header: %v", err)
	}
	bmpData.DIB.Width = int32(dibHeader[4]) | int32(dibHeader[5])<<8 | int32(dibHeader[6])<<16 | int32(dibHeader[7])<<24
	bmpData.DIB.Height = int32(dibHeader[8]) | int32(dibHeader[9])<<8 | int32(dibHeader[10])<<16 | int32(dibHeader[11])<<24
	bmpData.DIB.ImageSize = uint32(dibHeader[20]) | uint32(dibHeader[21])<<8 | uint32(dibHeader[22])<<16 | uint32(dibHeader[23])<<24

	return nil
}
