package main

import (
	"crud_book/internal/handlers"
	"crud_book/internal/services"
	"crud_book/internal/storage/postgres"
	"log"
)

func main() {
	migrator := postgres.NewMigrator("internal/storage/postgres/migrations")
	err := migrator.RunMigrations()
	if err != nil {
		log.Fatal("Migrations failed:", err)
	}

	storage := postgres.NewConnect()
	defer storage.Close()

	bookStorage := postgres.NewBookStorage(storage.DB())
	userStorage := postgres.NewUserStorage(storage.DB())

	bookService := services.NewBookService(bookStorage)
	userService := services.NewUserService(userStorage)

	handlers := handlers.NewHandlers(bookService, userService)
	router := handlers.SetupRouter()

	router.Run(":8080")
}
