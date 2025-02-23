package jsonrepository

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"hot-coffee/internal/core/entities"
)

// Errors
var (
	ErrInventoryItemDoesntExist   = errors.New("inventory item doesn't exist by provided id")
	ErrInventoryItemAlreadyExists = errors.New("item already exists in inventory")
)

// Singleton pattern
var inventoryRepositoryInstance *inventoryRepository

type inventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository() *inventoryRepository {
	if inventoryRepositoryInstance != nil {
		return inventoryRepositoryInstance
	}

	// Инициализация подключения к SQLite
	db, err := sql.Open("sqlite3", "hot-coffee.db")
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	inventoryRepositoryInstance = &inventoryRepository{
		db: db,
	}

	// Создание таблицы, если она не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS inventory (
			inventory_item_id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			quantity INTEGER NOT NULL,
			unit_id INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	return inventoryRepositoryInstance
}

func (r *inventoryRepository) Create(item entities.InventoryItem) error {
	// Проверка, существует ли уже элемент с таким ID
	existingItem, err := r.GetById(item.IngredientID)
	if err == nil && existingItem.IngredientID != "" {
		return ErrInventoryItemAlreadyExists
	}

	// SQL-запрос для добавления нового элемента
	query := `
		INSERT INTO inventory (name, quantity, unit_id, created_at)
		VALUES (?, ?, ?, ?)
	`
	_, err = r.db.Exec(query, item.Name, item.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) GetAll() ([]entities.InventoryItem, error) {
	// SQL-запрос для получения всех элементов
	query := `
		SELECT inventory_item_id, name, quantity,unit_id
		FROM inventory
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entities.InventoryItem
	for rows.Next() {
		var item entities.InventoryItem
		err := rows.Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *inventoryRepository) GetById(id string) (entities.InventoryItem, error) {
	// SQL-запрос для получения элемента по ID
	query := `
		SELECT inventory_item_id, name, quantity, unit_id, created_at
		FROM inventory
		WHERE inventory_item_id = ?
	`
	row := r.db.QueryRow(query, id)

	var item entities.InventoryItem
	err := row.Scan(&item.IngredientID, &item.Name, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.InventoryItem{}, ErrInventoryItemDoesntExist
		}
		return entities.InventoryItem{}, err
	}

	return item, nil
}

func (r *inventoryRepository) Update(id string, item entities.InventoryItem) error {
	// SQL-запрос для обновления элемента
	query := `
		UPDATE inventory
		SET name = ?, quantity = ?, unit_id = ?, created_at = ?
		WHERE inventory_item_id = ?
	`
	_, err := r.db.Exec(query, item.Name, item.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) Delete(id string) error {
	// SQL-запрос для удаления элемента
	query := `
		DELETE FROM inventory
		WHERE inventory_item_id = ?
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
