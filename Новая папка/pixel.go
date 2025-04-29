package main

type Pixel struct {
	B byte // Синий
	G byte // Зеленый
	R byte // Красный
}

// Функция для установки пикселя
func setPixel(pixelData []Pixel, width, x, y int, r, g, b byte) {
	pos := y*width + x
	pixelData[pos] = Pixel{B: b, G: g, R: r}
}

// Определение структуры для DIB заголовка

// Определение структуры для BMP заголовка
type BMPHeader struct {
	FileSize   uint32
	DataOffset uint32
}

// Определение структуры для BMP данных
