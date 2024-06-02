package storage

import (
	"database/sql"
	"log"
)

func (store *PostgressStore) Init() {
	createItemTable(store.db)
	createUserTable(store.db)
}

func createUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		HashedPassword TEXT
	);`

	result, err := db.Exec(query)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Creating users table: ", result)
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

	log.Println("Creating items table: ", result)
}
