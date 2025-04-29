package structure

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// BMPdata представляет структуру BMP-файла
type BMPdata struct {
	Header   BMPHeader
	Pixels   []Pixel
	Width    int
	Height   int
	RowBytes int
}

type BMPHeader struct {
	FileSize        uint32
	Reserved1       uint16
	Reserved2       uint16
	DataOffset      uint32
	HeaderSize      uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	SizeImage       uint32
	XPixelsPerMeter int32
	YPixelsPerMeter int32
	ColorsUsed      uint32
	ImportantColors uint32
}

type Pixel struct {
	B byte
	G byte
	R byte
}

// Функция для чтения BMP-файла
func ReadBMP(filename string) (*BMPdata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open BMP file: %w", err)
	}
	defer file.Close()

	// Чтение заголовка
	header := BMPHeader{}
	if err := binary.Read(file, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("failed to read BMP header: %w", err)
	}

	// Проверка формата
	if header.BitsPerPixel != 24 {
		return nil, errors.New("unsupported BMP format: only 24bpp is supported")
	}

	// Переход к пиксельным данным
	if _, err := file.Seek(int64(header.DataOffset), os.SEEK_SET); err != nil {
		return nil, fmt.Errorf("failed to seek to pixel data: %w", err)
	}

	width := int(header.Width)
	height := int(header.Height)
	rowBytes := ((width * 3) + 3) & ^3 // Учитываем выравнивание строк
	pixels := make([]Pixel, width*height)

	// Чтение пикселей
	for y := height - 1; y >= 0; y-- { // Пиксели в BMP идут снизу вверх
		row := make([]byte, rowBytes)
		if _, err := file.Read(row); err != nil {
			return nil, fmt.Errorf("failed to read BMP pixel row: %w", err)
		}
		for x := 0; x < width; x++ {
			i := y*width + x
			pixels[i] = Pixel{
				B: row[x*3],
				G: row[x*3+1],
				R: row[x*3+2],
			}
		}
	}

	return &BMPdata{
		Header:   header,
		Pixels:   pixels,
		Width:    width,
		Height:   height,
		RowBytes: rowBytes,
	}, nil
}
