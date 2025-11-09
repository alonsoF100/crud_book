package main

import (
	"crud_book/internal/handlers"
	"crud_book/internal/services"
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

	bookService := services.NewBookService(storage)
	userService := services.NewUserService(storage)

	handler := handlers.New(bookService, userService)
	router := handlers.SetupRouter(handler)

	router.Run(":8080")
}
