package s3

import (
	"log/slog"
	"os"

	"1337b04rd/internal/ports/external"
)

// NewImageStorageFromEnv создает новый экземпляр хранилища изображений из переменных окружения
func NewImageStorageFromEnv() external.ImageStorage {
	// Получаем URL хранилища из переменной окружения
	storageURL := os.Getenv("STORAGE_URL")
	if storageURL == "" {
		storageURL = "http://localhost:8082" // Значение по умолчанию
		slog.Info("STORAGE_URL не установлен, используется значение по умолчанию", "url", storageURL)
	}

	storage := NewS3ImageStorage(storageURL)

	// Создаем бакеты по умолчанию для постов и комментариев
	if err := storage.CreateBucketIfNotExists("posts"); err != nil {
		slog.Error("Ошибка при создании бакета для постов", "error", err)
	}

	if err := storage.CreateBucketIfNotExists("comments"); err != nil {
		slog.Error("Ошибка при создании бакета для комментариев", "error", err)
	}

	return storage
}
