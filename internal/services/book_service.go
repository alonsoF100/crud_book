package services

import (
	"crud_book/internal/models"
	"crud_book/internal/storage/postgres"
)

type BookService struct {
	storage *postgres.Storage
}

func NewBookService(storage *postgres.Storage) *BookService {
	return &BookService{storage: storage}
}

func (s *BookService) CreateBook(userID, name, description string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.CreateBook(userID, name, description)
}

func (s *BookService) GetUserBooks(userID string) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.GetUserBooks(userID)
}

func (s *BookService) UpdateBookStatus(bookID, status string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.storage.UpdateBookStatus(bookID, status)
	if err != nil {
		return nil, err
	}
	return s.storage.GetBook(bookID)
}

func (s *BookService) UpdateBookRating(bookID string, rating float64) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.storage.UpdateBookRating(bookID, rating)
	if err != nil {
		return nil, err
	}
	return s.storage.GetBook(bookID)
}

func (s *BookService) DeleteBook(bookID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.DeleteBook(bookID)
}

func (s *BookService) GetBook(bookID string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.GetBook(bookID)
}
