package api

import "github/bodzio50318/saleslist/storage"

type loginBody struct {
	UserName string
	Password string
}

type ItemsPage struct {
	Items []storage.Item
	itemFormData
}

func newItemsPage() *ItemsPage {
	return &ItemsPage{
		Items:        []storage.Item{},
		itemFormData: *newItemFormData(),
	}
}

type itemFormData struct {
	Values map[string]string
	Errors map[string]string
}

func newItemFormData() *itemFormData {
	return &itemFormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type jwtRespone struct {
	UserName string
	JwtToken string
}
