package handlers

import (
	"crud_book/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateUserHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте POST", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	user, err := storage.AddUser(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте GET", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID пользователя обязателен", http.StatusBadRequest)
		return
	}

	id := pathParts[2]

	user, err := storage.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте DELETE", r.Method), http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID пользователя обязателен", http.StatusBadRequest)
		return
	}

	userID := pathParts[2]

	err := storage.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllUsersHandler(storage *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Метод %s не разрешен. Используйте GET", r.Method), http.StatusMethodNotAllowed)
		return
	}

	users := storage.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
