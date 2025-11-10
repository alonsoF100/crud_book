package services

import (
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type BookService struct {
	repo storage.BookRepository
}

func NewBookService(repo storage.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(userID, name, description string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateBook(userID, name, description)
}

func (s *BookService) GetUserBooks(userID string) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUserBooks(userID)
}

func (s *BookService) UpdateBookStatus(bookID, status string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookStatus(bookID, status)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *BookService) UpdateBookRating(bookID string, rating float64) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookRating(bookID, rating)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *BookService) DeleteBook(bookID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteBook(bookID)
}

func (s *BookService) GetBook(bookID string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetBook(bookID)
}
