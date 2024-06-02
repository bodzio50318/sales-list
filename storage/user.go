package storage

import "log"

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
