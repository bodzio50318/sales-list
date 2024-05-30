package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "saleslist"
)

type Storage interface {
	Init()
	GetItems() ([]Item, error)
	// GetItemById(int) (*Item, error)
	// CreateItem(*Item) error
	// DeleteAccount(int) error
}

type PostgressStore struct {
	db *sql.DB
}

func (store *PostgressStore) GetItems() ([]Item, error) {
	query := "SELECT id, name FROM items"

	rows, err := store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item

	log.Println("Result: ", rows)

	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (store *PostgressStore) Init() {
	createItemTable(store.db)
}

func createItemTable(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS items (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );`

	result, err := db.Exec(query)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Creating item table: ", result)

}

func NewPostgressStore() *PostgressStore {
	user := os.Getenv("postgressUser")
	password := os.Getenv("postgressPassword")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println("User: ", user)
	log.Println("Password: ", password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	store := PostgressStore{db: db}
	store.Init()
	return &store
}
