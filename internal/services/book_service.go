package services

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type bookServiceImpl struct {
	repo storage.BookRepository
}

func NewBookService(repo storage.BookRepository) *bookServiceImpl {
	return &bookServiceImpl{repo: repo}
}

func (s *bookServiceImpl) CreateBook(userID, name, description string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateBook(userID, name, description)
}

func (s *bookServiceImpl) GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUserBooks(req)
}

func (s *bookServiceImpl) UpdateBookStatus(bookID, status string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookStatus(bookID, status)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *bookServiceImpl) UpdateBookRating(bookID string, rating float64) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookRating(bookID, rating)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(bookID)
}

func (s *bookServiceImpl) DeleteBook(bookID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteBook(bookID)
}

func (s *bookServiceImpl) GetBook(bookID string) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetBook(bookID)
}
