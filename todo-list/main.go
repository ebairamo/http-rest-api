package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	// Открываем соединение с SQLite (файл создастся автоматически)
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу users
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT,
		age INTEGER
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Вставляем несколько записей
	insertUser(db, "Алексей", 30)
	insertUser(db, "Марина", 25)
	insertUser(db, "Иван", 40)

	// Выводим всех пользователей
	printUsers(db)

	// Обновляем запись
	updateUser(db, 1, "Александр", 31)

	// Выводим обновленный список пользователей
	printUsers(db)

	// Удаляем запись
	deleteUser(db, 2)

	// Выводим итоговый список пользователей
	printUsers(db)
}

func insertUser(db *sql.DB, name string, age int) {
	_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		log.Fatal(err)
	}
}

func printUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id, age int
		var name string
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func updateUser(db *sql.DB, id int, name string, age int) {
	_, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", name, age, id)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteUser(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
