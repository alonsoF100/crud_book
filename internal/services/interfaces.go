package services

import (
	"context"
	"crud_book/internal/dto"
	"crud_book/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error)
	GetUser(ctx context.Context, req dto.GetUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type BookService interface {
	CreateBook(ctx context.Context, req dto.CreateBookRequest) (*models.Book, error)
	GetUserBooks(ctx context.Context, req dto.GetUserBooksRequest) ([]*models.Book, error)
	UpdateBookStatus(ctx context.Context, req dto.UpdateBookStatusRequest) (*models.Book, error)
	UpdateBookRating(ctx context.Context, req dto.UpdateBookRatingRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, req dto.DeleteBookRequest) error
	GetBook(ctx context.Context, req dto.GetBookRequest) (*models.Book, error)
}
