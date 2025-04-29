package external

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// S3ImageStorage реализует интерфейс хранилища изображений через Triple-S
type S3ImageStorage struct {
	baseURL    string
	httpClient *http.Client
	buckets    map[string]bool
}

// NewS3ImageStorage создает новый экземпляр хранилища изображений
func NewS3ImageStorage(baseURL string) *S3ImageStorage {
	// Проверяем, что URL заканчивается на "/"
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	return &S3ImageStorage{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		buckets: make(map[string]bool),
	}
}

// CreateBucketIfNotExists создает бакет, если он не существует
func (s *S3ImageStorage) CreateBucketIfNotExists(bucketName string) error {
	// Если бакет уже был создан ранее, не делаем запрос
	if _, exists := s.buckets[bucketName]; exists {
		return nil
	}

	// Формируем URL для создания бакета
	url := s.baseURL + bucketName

	// Создаем PUT-запрос
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		slog.Error("Ошибка создания запроса для бакета", "bucket", bucketName, "error", err)
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Выполняем запрос
	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("Ошибка выполнения запроса создания бакета", "bucket", bucketName, "error", err)
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusConflict {
		slog.Error("Неуспешный ответ при создании бакета", "bucket", bucketName, "status", resp.StatusCode)
		return fmt.Errorf("ошибка создания бакета: статус %d", resp.StatusCode)
	}

	// Отмечаем бакет как созданный
	s.buckets[bucketName] = true
	slog.Info("Бакет создан или уже существует", "bucket", bucketName)
	return nil
}

// UploadImage загружает изображение в хранилище
func (s *S3ImageStorage) UploadImage(ctx context.Context, bucketName, fileName string, imageData []byte) (string, error) {
	// Создаем бакет, если он еще не существует
	if err := s.CreateBucketIfNotExists(bucketName); err != nil {
		return "", err
	}

	// Формируем URL для загрузки файла
	url := s.baseURL + bucketName + "/" + fileName

	// Создаем PUT-запрос с данными изображения
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(imageData))
	if err != nil {
		slog.Error("Ошибка создания запроса для загрузки изображения", "bucket", bucketName, "file", fileName, "error", err)
		return "", fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Устанавливаем Content-Type
	req.Header.Set("Content-Type", "image/jpeg") // По умолчанию или определить динамически

	// Выполняем запрос
	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("Ошибка выполнения запроса загрузки изображения", "bucket", bucketName, "file", fileName, "error", err)
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		slog.Error("Неуспешный ответ при загрузке изображения", "bucket", bucketName, "file", fileName, "status", resp.StatusCode)
		return "", fmt.Errorf("ошибка загрузки изображения: статус %d", resp.StatusCode)
	}

	// Формируем URL для доступа к изображению
	imageURL := s.baseURL + bucketName + "/" + fileName
	slog.Info("Изображение успешно загружено", "bucket", bucketName, "file", fileName, "url", imageURL)
	return imageURL, nil
}

// UploadImageFromFile загружает изображение из файла
func (s *S3ImageStorage) UploadImageFromFile(ctx context.Context, bucketName, fileName string, reader io.Reader) (string, error) {
	// Читаем все содержимое файла
	imageData, err := io.ReadAll(reader)
	if err != nil {
		slog.Error("Ошибка чтения данных изображения", "error", err)
		return "", fmt.Errorf("ошибка чтения данных изображения: %w", err)
	}

	// Вызываем основной метод загрузки
	return s.UploadImage(ctx, bucketName, fileName, imageData)
}

// GetImage получает изображение из хранилища
func (s *S3ImageStorage) GetImage(ctx context.Context, bucketName, fileName string) ([]byte, error) {
	// Формируем URL для получения файла
	url := s.baseURL + bucketName + "/" + fileName

	// Создаем GET-запрос
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.Error("Ошибка создания запроса для получения изображения", "bucket", bucketName, "file", fileName, "error", err)
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Выполняем запрос
	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("Ошибка выполнения запроса получения изображения", "bucket", bucketName, "file", fileName, "error", err)
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			slog.Warn("Изображение не найдено", "bucket", bucketName, "file", fileName)
			return nil, fmt.Errorf("изображение не найдено")
		}

		slog.Error("Неуспешный ответ при получении изображения", "bucket", bucketName, "file", fileName, "status", resp.StatusCode)
		return nil, fmt.Errorf("ошибка получения изображения: статус %d", resp.StatusCode)
	}

	// Читаем данные изображения
	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Ошибка чтения данных изображения", "bucket", bucketName, "file", fileName, "error", err)
		return nil, fmt.Errorf("ошибка чтения данных изображения: %w", err)
	}

	slog.Info("Изображение успешно получено", "bucket", bucketName, "file", fileName, "size", len(imageData))
	return imageData, nil
}

// DeleteImage удаляет изображение из хранилища
func (s *S3ImageStorage) DeleteImage(ctx context.Context, bucketName, fileName string) error {
	// Формируем URL для удаления файла
	url := s.baseURL + bucketName + "/" + fileName

	// Создаем DELETE-запрос
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		slog.Error("Ошибка создания запроса для удаления изображения", "bucket", bucketName, "file", fileName, "error", err)
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Выполняем запрос
	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("Ошибка выполнения запроса удаления изображения", "bucket", bucketName, "file", fileName, "error", err)
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			slog.Warn("Изображение для удаления не найдено", "bucket", bucketName, "file", fileName)
			return nil // Не считаем отсутствие файла ошибкой при удалении
		}

		slog.Error("Неуспешный ответ при удалении изображения", "bucket", bucketName, "file", fileName, "status", resp.StatusCode)
		return fmt.Errorf("ошибка удаления изображения: статус %d", resp.StatusCode)
	}

	slog.Info("Изображение успешно удалено", "bucket", bucketName, "file", fileName)
	return nil
}
