package handlers

import (
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"

	"1337b04rd/internal/ports/external"
)

// ImageHandler обрабатывает HTTP запросы для изображений
type ImageHandler struct {
	imageStorage external.ImageStorage
}

// NewImageHandler создает новый обработчик изображений
func NewImageHandler(imageStorage external.ImageStorage) *ImageHandler {
	return &ImageHandler{
		imageStorage: imageStorage,
	}
}

// HandleGetImage обрабатывает GET запрос для получения изображения
func (h *ImageHandler) HandleGetImage(w http.ResponseWriter, r *http.Request) {
	// Парсим путь запроса для получения bucket и fileName
	// Пример: /images/posts/image.jpg -> bucket=posts, fileName=image.jpg
	path := strings.TrimPrefix(r.URL.Path, "/images/")
	parts := strings.SplitN(path, "/", 2)

	if len(parts) != 2 {
		slog.Error("Неверный формат пути к изображению", "path", r.URL.Path)
		http.Error(w, "Неверный формат пути к изображению", http.StatusBadRequest)
		return
	}

	bucket := parts[0]
	fileName := parts[1]

	// Получаем изображение из хранилища
	imageData, err := h.imageStorage.GetImage(r.Context(), bucket, fileName)
	if err != nil {
		slog.Error("Ошибка получения изображения", "bucket", bucket, "fileName", fileName, "error", err)
		http.Error(w, "Изображение не найдено", http.StatusNotFound)
		return
	}

	// Определяем MIME-тип на основе расширения файла
	contentType := getContentTypeByExtension(fileName)

	// Устанавливаем заголовки и отправляем изображение
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", string(len(imageData)))
	w.Header().Set("Cache-Control", "public, max-age=86400") // Кеширование на 1 день
	w.WriteHeader(http.StatusOK)
	w.Write(imageData)
}

// getContentTypeByExtension определяет MIME-тип на основе расширения файла
func getContentTypeByExtension(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
