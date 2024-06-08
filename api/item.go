package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *ApiServer) handleGetItems(c echo.Context) error {
	items, err := s.store.GetItems()
	if err != nil {
		return err
	}

	log.Println("items", items)
	return c.Render(http.StatusOK, "itemPage", items)
}

func (s *ApiServer) handlePostItems(c echo.Context) error {
	name := c.FormValue("name")
	err := s.store.InsertItem(name)

	if err != nil {
		formData := newItemFormData()
		formData.Errors["name"] = "Failed to insert item"
		formData.Values["name"] = name
		return c.Render(http.StatusBadRequest, "itemForm", formData)
	}

	items, err := s.store.GetItems()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "listOfItems", items)
}

func (s *ApiServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *ApiServer) handleDeleteItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
