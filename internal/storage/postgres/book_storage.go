package postgres

import (
	"context"
	"crud_book/internal/dto"
	"crud_book/internal/models"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookStorage struct {
	db *pgxpool.Pool
}

func NewBookStorage(db *pgxpool.Pool) *bookStorage {
	return &bookStorage{db: db}
}

func (s *bookStorage) CreateBook(ctx context.Context, req dto.CreateBookRequest) (*models.Book, error) {
	newID := uuid.New().String()

	const query = `INSERT INTO books (id, user_id, name, description)
	               VALUES ($1, $2, $3, $4)
				   RETURNING id, user_id, name, description, rating, status, created_at`
				   
	var book models.Book
	err := s.db.QueryRow(ctx, query, newID, req.UserID, req.Name, req.Description).
		Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status, &book.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("create book: %w", err)
	}

	return &book, nil
}

func (s *bookStorage) UpdateBookStatus(ctx context.Context, req dto.UpdateBookStatusRequest) error {
	const query = `UPDATE books
	               SET status = $1
	               WHERE id = $2`

	result, err := s.db.Exec(ctx, query, req.Status, req.BookID)
	if err != nil {
		return fmt.Errorf("update book status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to find book with ID: %s", req.BookID)
	}

	return nil
}

func (s *bookStorage) UpdateBookRating(ctx context.Context, req dto.UpdateBookRatingRequest) error {
	const query = `UPDATE books
	               SET rating = $1
				   WHERE id = $2`

	result, err := s.db.Exec(ctx, query, req.Rating, req.BookID)
	if err != nil {
		return fmt.Errorf("update book rating: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to find book with ID: %s", req.BookID)
	}

	return nil
}

func (s *bookStorage) DeleteBook(ctx context.Context, req dto.DeleteBookRequest) error {
	const query = `DELETE FROM books
	               WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.BookID)
	if err != nil {
		return fmt.Errorf("delete book: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed find book with ID: %s", req.BookID)
	}

	return nil
}

func (s *bookStorage) GetBook(ctx context.Context, req dto.GetBookRequest) (*models.Book, error) {
	const query = `SELECT id, user_id, name, description, rating, status, created_at
	               FROM books WHERE id = $1`

	var book models.Book
	err := s.db.QueryRow(ctx, query, req.BookID).Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status, &book.CreatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("failed to find book with ID: %s", req.BookID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &book, nil
}

func (s *bookStorage) GetUserBooks(ctx context.Context, req dto.GetUserBooksRequest) ([]*models.Book, error) {
	qb := squirrel.
		Select("id", "user_id", "name", "description", "rating", "status", "created_at").
		From("books").
		Where(squirrel.Eq{"user_id": req.UserID})

	if req.Status != "" {
		qb = qb.Where(squirrel.Eq{"status": req.Status})
	}

	if req.MinRating != 0 || req.MaxRating != 0 {
		if req.MinRating != 0 {
			qb = qb.Where(squirrel.GtOrEq{"rating": req.MinRating})
		}
		if req.MaxRating != 0 {
			qb = qb.Where(squirrel.LtOrEq{"rating": req.MaxRating})
		}
	}

	query, args, err := qb.OrderBy(req.Sort + " " + req.Order).
		Limit(uint64(req.Limit)).
		Offset(uint64(req.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get users books: %w", err)
	}
	defer rows.Close()

	var userBooks []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status, &book.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}

		userBooks = append(userBooks, &book)
	}

	return userBooks, nil
}
