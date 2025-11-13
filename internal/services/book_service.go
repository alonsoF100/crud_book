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

func (s *bookServiceImpl) CreateBook(req dto.CreateBookRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateBook(req)
}

func (s *bookServiceImpl) GetUserBooks(req dto.GetUserBooksRequest) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUserBooks(req)
}

func (s *bookServiceImpl) UpdateBookStatus(req dto.UpdateBookStatusRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookStatus(req)
	if err != nil {
		return nil, err
	}
	
	return s.repo.GetBook(dto.NewGetBookRequest(req.BookID))
}

func (s *bookServiceImpl) UpdateBookRating(req dto.UpdateBookRatingRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookRating(req)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(dto.NewGetBookRequest(req.BookID))
}

func (s *bookServiceImpl) DeleteBook(req dto.DeleteBookRequest) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteBook(req)
}

func (s *bookServiceImpl) GetBook(req dto.GetBookRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetBook(req)
}
