package main

func applyBlueFilter(pixelData []Pixel, width, height, x, y int) {
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
func applyGreenFilter(pixelData []Pixel, width, height, x, y int) {
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
func applyRedFilter(pixelData []Pixel) {
	for i := range pixelData {
		pixelData[i] = Pixel{
			B: 0,
			G: 0,
			R: pixelData[i].R,
		}
	}
}
func applyNegativeFilter(pixelData []Pixel, width, height, x, y int) {
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
func applyBlurFilter(pixelData []Pixel, width, height, radius int) {
	// Создаем новый массив для хранения размытых пикселей
	newPixelData := make([]Pixel, len(pixelData))

	// Применяем фильтр размытия к каждому пикселю
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var sumR, sumG, sumB int // Сумма цветов для усреднения
			count := 0               // Счётчик пикселей в области размытия

			// Проходим по окрестности пикселя, используя радиус размытия
			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					ni := y + dy // Индекс строки соседнего пикселя
					nj := x + dx // Индекс столбца соседнего пикселя

					// Проверка, что соседний пиксель находится в пределах изображения
					if ni >= 0 && ni < height && nj >= 0 && nj < width {
						pixel := pixelData[ni*width+nj] // Получаем цвет соседнего пикселя
						sumR += int(pixel.R)
						sumG += int(pixel.G)
						sumB += int(pixel.B)
						count++ // Увеличиваем счётчик пикселей
					}
				}
			}

			// Если были найдены соседние пиксели, усредняем их цвет
			if count > 0 {
				newPixelData[y*width+x] = Pixel{
					R: byte(sumR / count),
					G: byte(sumG / count),
					B: byte(sumB / count),
				}
			} else {
				newPixelData[y*width+x] = pixelData[y*width+x] // Если соседей нет, оставляем исходный цвет
			}
		}
	}

	// Обновляем оригинальные данные пикселей
	copy(pixelData, newPixelData)
}

func applyPixelateFilter(pixelData []Pixel, width, height, blockSize int) []Pixel {
	for y := 0; y < height; y += blockSize {
		for x := 0; x < width; x += blockSize {
			// Найти центр блока
			centerX := x + blockSize/2
			centerY := y + blockSize/2

			// Убедиться, что центр не выходит за границы изображения
			if centerX >= width {
				centerX = width - 1
			}
			if centerY >= height {
				centerY = height - 1
			}

			// Получить цвет пикселя в центре
			centerPixel := pixelData[centerY*width+centerX]

			// Перекрасить весь блок в цвет центрального пикселя
			for dy := 0; dy < blockSize && (y+dy) < height; dy++ {
				for dx := 0; dx < blockSize && (x+dx) < width; dx++ {
					pixelData[(y+dy)*width+(x+dx)] = centerPixel
				}
			}
		}
	}
	return pixelData
}
