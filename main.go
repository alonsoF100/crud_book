package main

import (
	"crud_book/handlers"
	"crud_book/storage"
	"log"
)

func main() {
	storage := storage.NewConnect()
	defer storage.Close()

	err := storage.RunMigrations()
	if err != nil {
		log.Fatal("Migrations failed:", err)
	}

	handler := handlers.New(storage)
	router := handlers.SetupRouter(handler)

	router.Run(":8080")
}
