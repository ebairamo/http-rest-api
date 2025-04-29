package cmd

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	httpAdapter "1337b04rd/internal/adapters/primary/http"
	"1337b04rd/internal/adapters/secondary/postgres"
)

func Run() {
	// Используем обычный log пакет для гарантированного вывода
	log.Println("Запуск программы...")

	// Парсинг командной строки
	port := flag.Int("port", 8080, "Номер порта")
	help := flag.Bool("help", false, "Показать справку")
	flag.Parse()

	log.Println("Параметры обработаны, порт:", *port, "справка:", *help)

	// Показать помощь и выйти, если запрошена помощь
	if *help {
		fmt.Println("hacker board")
		fmt.Println("\nИспользование:")
		fmt.Println("  1337b04rd [--port <N>]")
		fmt.Println("  1337b04rd --help")
		fmt.Println("\nОпции:")
		fmt.Println("  --help       Показать эту справку.")
		fmt.Println("  --port N     Номер порта.")
		os.Exit(0)
	}

	// Настройка логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// Проверка доступности сервиса хранилища
	storageURL := os.Getenv("STORAGE_URL")
	if storageURL == "" {
		storageURL = "http://localhost:8082" // Значение по умолчанию
	}

	// Проверяем соединение с сервисом хранилища
	client := http.Client{Timeout: 5 * time.Second}
	_, err := client.Get(storageURL)
	if err != nil {
		slog.Warn("Сервис хранилища изображений недоступен", "url", storageURL, "error", err)
		slog.Info("Продолжаем без хранилища изображений, загрузка файлов может не работать")
	} else {
		slog.Info("Сервис хранилища изображений доступен", "url", storageURL)
	}

	// Подключение к базе данных
	db := postgres.Connect()
	defer db.Close()
	logger.Info("База данных подключена")

	// Запускаем HTTP сервер с передачей порта и соединения с БД
	httpAdapter.StartServer(*port, db)
}
