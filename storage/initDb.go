package storage

import (
	"database/sql"
	"fmt"
	"log"
)

func (store *PostgressStore) Init() {
	createItemTable(store.db)
	createUserTable(store.db)

	users, err := store.GetUserList()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Users: ", users)

	if len(users) == 0 {
		fmt.Println("Inserting admin user")
		store.InsertUser("bodzio")
	}

	items, err := store.GetItems()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Items: ", items)
	if len(items) == 0 {
		fmt.Println("Inserting items")
		store.InsertItem("jajka")
		store.InsertItem("kawa")
		store.InsertItem("koks")
	}

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
            name TEXT NOT NULL UNIQUE
        );`

	result, err := db.Exec(query)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Creating items table: ", result)
}
