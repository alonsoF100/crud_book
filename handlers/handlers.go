package handlers

import (
	"crud_book/storage"
	"net/http"
	"strings"
)

func SetupRoutes(s *storage.Storage) {
	// Обработка /users (GET и POST)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			CreateUserHandler(s, w, r) // POST /users - создание пользователя
		case "GET":
			GetAllUsersHandler(s, w, r) // GET /users - все пользователи
		default:
			http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		}
	})

	// Обработка /books (только POST)
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			CreateBookHandler(s, w, r) // POST /books - создание книги
		} else {
			http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		}
	})

	// Обработка /users/...
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")

		switch {
		case len(pathParts) == 4 && pathParts[3] == "books":
			GetUserBooksHandler(s, w, r) // GET /users/:id/books - книги пользователя
		case len(pathParts) == 3 && r.Method == "GET":
			GetUserHandler(s, w, r) // GET /users/:id - получение пользователя
		case len(pathParts) == 3 && r.Method == "DELETE":
			DeleteUserHandler(s, w, r) // DELETE /users/:id - удаление пользователя
		default:
			http.NotFound(w, r)
		}
	})

	// Обработка /books/...
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")

		switch {
		case len(pathParts) == 4 && pathParts[3] == "status":
			UpdateBookStatusHandler(s, w, r) // PUT /books/:id/status - обновление статуса
		case len(pathParts) == 4 && pathParts[3] == "rating":
			UpdateBookRatingHandler(s, w, r) // PUT /books/:id/rating - обновление рейтинга
		case len(pathParts) == 3 && r.Method == "GET":
			GetBookHandler(s, w, r) // GET /books/:id - получение книги
		case len(pathParts) == 3 && r.Method == "DELETE":
			DeleteBookHandler(s, w, r) // DELETE /books/:id - удаление книги
		default:
			http.NotFound(w, r)
		}
	})
}
