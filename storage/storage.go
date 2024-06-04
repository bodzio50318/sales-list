package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
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

var DB_HOST = os.Getenv("DB_HOST")
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_PORT = os.Getenv("DB_PORT")
var DB_NAME = os.Getenv("DB_NAME")

func NewPostgressStore() *PostgressStore {

	log.Println("Db Host is: ", DB_HOST)

	intPort, err := strconv.Atoi(DB_PORT)

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, intPort, DB_USER, DB_PASSWORD, DB_NAME)

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
