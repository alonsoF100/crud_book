package main

import (
	"crud_book/internal/handlers"
	"crud_book/internal/storage/postgres"
	"log"
)

func main() {
	storage := postgres.NewConnect()
	defer storage.Close()

	err := storage.RunMigrations()
	if err != nil {
		log.Fatal("Migrations failed:", err)
	}

	handler := handlers.New(storage)
	router := handlers.SetupRouter(handler)

	router.Run(":8080")
}
