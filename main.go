package main

import (
	"github/bodzio50318/saleslist/api"
	"github/bodzio50318/saleslist/storage"
	"log"
	"os"
)

func main() {
	storage := storage.NewPostgressStore()
	path := os.Getenv("salesPath")
	log.Println("Os path:", path)
	if path == "" {
		path = "0.0.0.0:8080"
	}

	server := api.NewApiServer(path, storage)
	server.Run()
	log.Println("Server started good!")
}
