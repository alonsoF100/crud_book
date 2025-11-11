package services

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
)

type UserService interface {
	CreateUser(name, email string) (*models.User, error)
	GetUser(userID string) (*models.User, error)
	DeleteUser(userID string) error
	GetAllUsers() ([]*models.User, error)
}

type BookService interface {
	CreateBook(userID, name, description string) (*models.Book, error)
	GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error)
	UpdateBookStatus(bookID, status string) (*models.Book, error)
	UpdateBookRating(bookID string, rating float64) (*models.Book, error)
	DeleteBook(bookID string) error
	GetBook(bookID string) (*models.Book, error)
}
