package storage

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
)

type BookRepository interface {
	CreateBook(userID, name, description string) (*models.Book, error)
	GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error)
	UpdateBookStatus(bookID, status string) error
	UpdateBookRating(bookID string, rating float64) error
	GetBook(bookID string) (*models.Book, error)
	DeleteBook(bookID string) error
}

type UserRepository interface {
	CreateUser(name, email string) (*models.User, error)
	GetUser(userID string) (*models.User, error)
	DeleteUser(userID string) error
	GetAllUsers() ([]*models.User, error)
}
