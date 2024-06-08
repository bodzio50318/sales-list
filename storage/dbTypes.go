package storage

import "database/sql"

type Items []Item

type Item struct {
	Id   int
	Name string
}

type User struct {
	Id             int
	Name           string
	HashedPassword sql.NullString
}