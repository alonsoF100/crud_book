package services

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*models.User, error)
	GetUser(req dto.GetUserRequest) (*models.User, error)
	DeleteUser(req dto.DeleteUserRequest) error
	GetAllUsers() ([]*models.User, error)
}

type BookService interface {
	CreateBook(req dto.CreateBookRequest) (*models.Book, error)
	GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error)
	UpdateBookStatus(req dto.UpdateBookStatusRequest) (*models.Book, error)
	UpdateBookRating(req dto.UpdateBookRatingRequest) (*models.Book, error)
	DeleteBook(req dto.DeleteBookRequest) error
	GetBook(req dto.GetBookRequest) (*models.Book, error)
}
