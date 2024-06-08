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
	itemsPageData := newItemsPageData()

	itemsPageData.Items = items
	return c.Render(http.StatusOK, "itemPage", itemsPageData)
}

func (s *ApiServer) handlePostItems(c echo.Context) error {

	name := c.FormValue("name")

	formData := newItemFormData()
	formData.Values["name"] = name

	item, err := s.store.InsertItem(name)

	if err != nil {
		log.Println("Failed to insert item", err)
		formData.Errors["name"] = "Failed to insert item"
		return c.Render(http.StatusUnprocessableEntity, "itemForm", formData)
	}
	c.Render(http.StatusOK, "itemForm", newItemFormData())
	return c.Render(http.StatusOK, "oob-item", item)
}

func (s *ApiServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *ApiServer) handleDeleteItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
