package handlers

import (
	"crud_book/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateBookHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте POST", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var req CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	book, err := storage.AddBook(req.UserID, req.Name, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetUserBooksHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте GET", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 4 || pathParts[2] == "" {
		http.Error(w, "ID пользователя обязателен", http.StatusBadRequest)
		return
	}

	userID := pathParts[2]

	books, err := storage.GetUserBooks(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func UpdateBookStatusHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте PUT", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 4 || pathParts[2] == "" {
		http.Error(w, "ID книги обязателен", http.StatusBadRequest)
		return
	}

	bookID := pathParts[2]

	var req UpdateBookStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	err := storage.UpdateBookStatus(bookID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func UpdateBookRatingHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте PUT", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 4 || pathParts[2] == "" {
		http.Error(w, "ID книги обязателен", http.StatusBadRequest)
		return
	}

	bookID := pathParts[2]

	var req UpdateBookRatingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	err := storage.UpdateBookRating(bookID, req.Rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "rating updated"})
}

func DeleteBookHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте DELETE", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID книги обязателен", http.StatusBadRequest)
		return
	}

	bookID := pathParts[2]

	err := storage.DeleteBook(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetBookHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте GET", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID книги обязателен", http.StatusBadRequest)
		return
	}

	bookID := pathParts[2]

	book, err := storage.GetBook(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
