package storage

import (
	"fmt"
	"log"
)

func (store *PostgressStore) SetUserPassword(id int, hashedPassword string) error {
	query := "UPDATE users SET hashedpassword = $1 WHERE id = $2"
	_, err := store.db.Exec(query, string(hashedPassword), id)
	if err != nil {
		log.Printf("Error setting user password: %v", err)
		return err
	}
	return nil
}

func (store *PostgressStore) GetUserByName(name string) (*User, error) {
	query := "SELECT * FROM users WHERE name = $1"

	log.Println(name)
	row := store.db.QueryRow(query, name)

	var user User
	if err := row.Scan(&user.Id, &user.Name, &user.HashedPassword); err != nil {
		return nil, err
	}

	return &user, nil
}

func (store *PostgressStore) InsertUser(name string) error {
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"

	row := store.db.QueryRow(query, name)

	fmt.Println("Row: ", row)
	return nil
}

func (store *PostgressStore) GetUserList() ([]User, error) {
	query := "SELECT id, name FROM users"

	rows, err := store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
