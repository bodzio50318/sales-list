package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github/bodzio50318/saleslist/storage"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiServer struct {
	listenAddress string
	store         storage.Storage
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/item", makeHttpHandleFunc(s.handleItem))
	router.HandleFunc("/login", makeHttpHandleFunc(s.handleLogin))

	log.Println("Starting a server on port: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err.Error())
			WriteJson(w, http.StatusBadRequest, err.Error())
		}
	}
}

func NewApiServer(listenAddress string, store storage.Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
