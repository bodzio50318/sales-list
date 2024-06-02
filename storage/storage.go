package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	port   = 5432
	dbname = "saleslist"
)

type Storage interface {
	Init()
	GetItems() ([]Item, error)
	GetUserByName(string) (*User, error)
	SetUserPassword(int, string) error
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

func NewPostgressStore() *PostgressStore {
	host := os.Getenv("db_host")

	log.Println("Db Host is: ", host)
	user := os.Getenv("postgressUser")
	password := os.Getenv("postgressPassword")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
