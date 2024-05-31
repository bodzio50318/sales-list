package main

import (
	"github/bodzio50318/saleslist/api"
	"github/bodzio50318/saleslist/storage"
	"log"
)

func main() {
	storage := storage.NewPostgressStore()
	server := api.NewApiServer("0.0.0.0:8080", storage)
	server.Run()
	log.Println("Server started good!")
}
