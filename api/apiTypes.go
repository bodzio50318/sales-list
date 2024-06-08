package api

import "github/bodzio50318/saleslist/storage"

type LoginBody struct {
	UserName string
	Password string
}
type ItemsPageData struct {
	Items    storage.Items
	FormData ItemFormData
}
type ItemFormData struct {
	Values map[string]string
	Errors map[string]string
}
type JwtRespone struct {
	UserName string
	JwtToken string
}

func newItemsPageData() *ItemsPageData {
	return &ItemsPageData{
		Items:    make(storage.Items, 0),
		FormData: *newItemFormData(),
	}
}

func newItemFormData() *ItemFormData {
	return &ItemFormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
