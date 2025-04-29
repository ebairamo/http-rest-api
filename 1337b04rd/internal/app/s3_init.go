package cmd

import (
	"context"
	"log/slog"
	"time"

	"1337b04rd/internal/ports/external"
)

const (
	PostImageBucket    = "post-images"
	CommentImageBucket = "comment-images"
	AvatarBucket       = "avatars"

	// Максимальное количество попыток
	maxInitRetries = 5
	retryInterval  = 2 * time.Second
)

// InitializeS3Buckets инициализирует бакеты для изображений
func InitializeS3Buckets(storage external.ImageStorage) error {
	ctx := context.Background()

	// Создаем бакеты с ретраями
	buckets := []string{PostImageBucket, CommentImageBucket, AvatarBucket}

	for _, bucket := range buckets {
		if err := createBucketWithRetry(ctx, storage, bucket); err != nil {
			return err
		}
	}

	slog.Info("S3 бакеты успешно инициализированы")
	return nil
}

// createBucketWithRetry создает бакет с повторными попытками
func createBucketWithRetry(ctx context.Context, storage external.ImageStorage, bucketName string) error {
	var err error

	for i := 0; i < maxInitRetries; i++ {
		err = storage.CreateBucket(ctx, bucketName)
		if err == nil {
			slog.Info("Бакет успешно создан", "bucket", bucketName)
			return nil
		}

		slog.Warn("Не удалось создать бакет, повторная попытка",
			"bucket", bucketName,
			"attempt", i+1,
			"maxAttempts", maxInitRetries,
			"error", err)

		time.Sleep(retryInterval)
	}

	slog.Error("Не удалось создать бакет после нескольких попыток",
		"bucket", bucketName,
		"attempts", maxInitRetries,
		"error", err)

	return err
}
