package utils

import (
	strct "bitmap/structure"
	"encoding/binary"
	"fmt"
	"os"
)

// Печать заголовка BMP-файла
func ReadBMPHeader(filename string) error {
	data, err := strct.ReadBMP(filename)
	if err != nil {
		return err
	}
	header := data.Header
	fmt.Printf("File Size: %d\n", header.FileSize)
	fmt.Printf("Image Width: %d\n", header.Width)
	fmt.Printf("Image Height: %d\n", header.Height)
	fmt.Printf("Bits Per Pixel: %d\n", header.BitsPerPixel)
	return nil
}

// Сохранение BMP-файла
func SaveBMP(data *strct.BMPdata, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Запись заголовка
	if err := binary.Write(file, binary.LittleEndian, &data.Header); err != nil {
		return fmt.Errorf("failed to write BMP header: %w", err)
	}

	// Запись пиксельных данных
	rowBytes := data.RowBytes
	for y := data.Height - 1; y >= 0; y-- {
		row := make([]byte, rowBytes)
		for x := 0; x < data.Width; x++ {
			pixel := data.Pixels[y*data.Width+x]
			row[x*3] = pixel.B
			row[x*3+1] = pixel.G
			row[x*3+2] = pixel.R
		}
		if _, err := file.Write(row); err != nil {
			return fmt.Errorf("failed to write pixel data: %w", err)
		}
	}
	return nil
}

// Поворот изображения
func RotateImage(data *strct.BMPdata, rotations []string) error {
	// Реализация поворота изображения
	// Нужно учитывать угол поворота
	// Возвращаем новую ширину и высоту, если поворот на 90 или 270 градусов
	return nil
}

// Обрезка изображения
func CropImage(data *strct.BMPdata, crops []string) error {
	// Реализация обрезки изображения
	return nil
}

// Отражение изображения
func MirrorHorizontal(data *strct.BMPdata) {
	width := data.Width
	height := data.Height
	for y := 0; y < height; y++ {
		for x := 0; x < width/2; x++ {
			i1 := y*width + x
			i2 := y*width + (width - x - 1)
			data.Pixels[i1], data.Pixels[i2] = data.Pixels[i2], data.Pixels[i1]
		}
	}
}

func MirrorVertical(data *strct.BMPdata) {
	width := data.Width
	height := data.Height
	for y := 0; y < height/2; y++ {
		for x := 0; x < width; x++ {
			i1 := y*width + x
			i2 := (height-1-y)*width + x
			data.Pixels[i1], data.Pixels[i2] = data.Pixels[i2], data.Pixels[i1]
		}
	}
}

// Основная справка
func MainHelp() string {
	return `
Usage:
  bitmap <command> [options]

Available Commands:
  header   Prints the BMP file header.
  apply    Applies transformations to BMP file (filters, rotation, etc.).
  help     Show this help message.
`
}

// Справка по заголовку
func HelpHeader() string {
	return `
Usage:
  bitmap header <inputfile>

Description:
  This command prints out the header information of a BMP file.
`
}

// Справка по применению фильтров
func HelpApply() string {
	return `
Usage:
  bitmap apply <inputfile> <outputfile> [options]

Options:
  -filter    Apply a filter (blue, red, green, negative, grayscale)
  -rotate    Rotate the image (90, 180, 270 degrees)
  -crop      Crop the image with the given parameters (x, y, width, height)
  -mirror    Mirror the image (horizontal, vertical)
`
}
