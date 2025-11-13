package services

import (
	"context"
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

func (s *bookServiceImpl) CreateBook(ctx context.Context, req dto.CreateBookRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateBook(ctx, req)
}

func (s *bookServiceImpl) GetUserBooks(ctx context.Context, req dto.GetUserBooksRequest) ([]*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUserBooks(ctx, req)
}

func (s *bookServiceImpl) UpdateBookStatus(ctx context.Context, req dto.UpdateBookStatusRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return s.repo.GetBook(ctx, dto.NewGetBookRequest(req.BookID))
}

func (s *bookServiceImpl) UpdateBookRating(ctx context.Context, req dto.UpdateBookRatingRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	err := s.repo.UpdateBookRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBook(ctx, dto.NewGetBookRequest(req.BookID))
}

func (s *bookServiceImpl) DeleteBook(ctx context.Context, req dto.DeleteBookRequest) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteBook(ctx, req)
}

func (s *bookServiceImpl) GetBook(ctx context.Context, req dto.GetBookRequest) (*models.Book, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetBook(ctx, req)
}
