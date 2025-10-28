package main

import (
	"crud_book/handlers"
	"crud_book/storage"
	"net/http"
)

func main() {
	storage := storage.NewStorage()
	handlers.SetupRoutes(storage)

	http.ListenAndServe(":8080", nil)
}
