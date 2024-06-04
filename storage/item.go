package storage

import "fmt"

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

func (store *PostgressStore) InsertItem(itemName string) error {
	query := "INSERT INTO items (name) VALUES ($1) RETURNING id"

	row := store.db.QueryRow(query, itemName)

	fmt.Println("Row: ", row)
	return nil
}
