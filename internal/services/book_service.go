package services

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type BookServiceImpl struct {
	repo storage.BookRepository
}

func NewBookService(repo storage.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{repo: repo}
}

func (s *BookServiceImpl) CreateBook(userID, name, description string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateBook(userID, name, description)
}

func (s *BookServiceImpl) GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUserBooks(req)
}

func (s *BookServiceImpl) UpdateBookStatus(bookID, status string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookStatus(bookID, status)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *BookServiceImpl) UpdateBookRating(bookID string, rating float64) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookRating(bookID, rating)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *BookServiceImpl) DeleteBook(bookID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteBook(bookID)
}

func (s *BookServiceImpl) GetBook(bookID string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetBook(bookID)
}
