package filter

import (
	strct "bitmap/structure"
)

// Применение фильтра "синий"
func FilterBlue(data *strct.BMPdata) {
	for i := range data.Pixels {
		data.Pixels[i].R = 0
		data.Pixels[i].G = 0
	}
}

// Применение фильтра "красный"
func FilterRed(data *strct.BMPdata) {
	for i := range data.Pixels {
		data.Pixels[i].G = 0
		data.Pixels[i].B = 0
	}
}

// Применение фильтра "зеленый"
func FilterGreen(data *strct.BMPdata) {
	for i := range data.Pixels {
		data.Pixels[i].R = 0
		data.Pixels[i].B = 0
	}
}

// Фильтр "негатив"
func FilterNegative(data *strct.BMPdata) {
	for i := range data.Pixels {
		data.Pixels[i].R = 255 - data.Pixels[i].R
		data.Pixels[i].G = 255 - data.Pixels[i].G
		data.Pixels[i].B = 255 - data.Pixels[i].B
	}
}

// Фильтр "черно-белый"
func FilterGrayScale(data *strct.BMPdata) {
	for i := range data.Pixels {
		gray := (uint16(data.Pixels[i].R) + uint16(data.Pixels[i].G) + uint16(data.Pixels[i].B)) / 3
		data.Pixels[i].R = byte(gray)
		data.Pixels[i].G = byte(gray)
		data.Pixels[i].B = byte(gray)
	}
}
