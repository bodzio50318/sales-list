package api

import (
	"fmt"
	"net/http"
)

func (s *ApiServer) handleItem(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetItems(w, r)
	}
	return fmt.Errorf("method not supported")
}

func (s *ApiServer) handleGetItems(w http.ResponseWriter, r *http.Request) error {
	result, err := s.store.GetItems()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, result)
}
func (s *ApiServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *ApiServer) handleDeleteItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
