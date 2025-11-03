package main

import (
	"crud_book/handlers"
	"crud_book/storage"
)

func main() {
	storage := storage.NewConnect()
	defer storage.Close()

	handler := handlers.New(storage)
	router := handlers.SetupRouter(handler)

	router.Run(":8080")
}
