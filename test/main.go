package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Структура для ведра
type Bucket struct {
	Name             string `json:"Name"`
	CreationTime     string `json:"CreationTime"`
	LastModifiedTime string `json:"LastModifiedTime"`
	Status           string `json:"Status"`
}

// Обработчик для списка ведер
func ListBucketsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	buckets := []Bucket{
		{Name: "bucket1", CreationTime: "2024-10-16T10:00:00Z", LastModifiedTime: "2024-10-17T12:00:00Z", Status: "active"},
		{Name: "bucket2", CreationTime: "2024-10-15T09:00:00Z", LastModifiedTime: "2024-10-16T14:00:00Z", Status: "active"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buckets)
}

// Обработчик для получения объекта
func GetObjectHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	bucketName := parts[1]
	objectKey := parts[2]

	// Пример данных объекта
	objectData := "This is the object data."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bucket": bucketName,
		"object": objectKey,
		"data":   objectData,
	})
}

func main() {
	http.HandleFunc("/buckets", ListBucketsHandler)
	http.HandleFunc("/buckets/", GetObjectHandler)

	http.ListenAndServe(":8080", nil)
}
