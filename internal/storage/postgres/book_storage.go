package postgres

import (
	"context"
	"crud_book/internal/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookStorage struct {
	db *pgxpool.Pool
}

func NewBookStorage(db *pgxpool.Pool) *BookStorage {
	return &BookStorage{db: db}
}

func (s *BookStorage) CreateBook(userID, name, description string) (*models.Book, error) {
	newID := uuid.New().String()

	const query = `INSERT INTO books (id, user_id, name, description)
	               VALUES ($1, $2, $3, $4)
				   RETURNING id, user_id, name, description, rating, status`
	var book models.Book
	err := s.db.QueryRow(context.Background(), query, newID, userID, name, description).
		Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
	if err != nil {
		return nil, fmt.Errorf("create book: %w", err)
	}

	return &book, nil
}

func (s *BookStorage) UpdateBookStatus(bookID, status string) error {
	const query = `UPDATE books
	               SET status = $1
	               WHERE id = $2`

	result, err := s.db.Exec(context.Background(), query, status, bookID)
	if err != nil {
		return fmt.Errorf("update book status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to find book with ID: %s", bookID)
	}

	return nil
}

func (s *BookStorage) UpdateBookRating(bookID string, rating float64) error {
	const query = `UPDATE books
	               SET rating = $1
				   WHERE id = $2`

	result, err := s.db.Exec(context.Background(), query, rating, bookID)
	if err != nil {
		return fmt.Errorf("update book rating: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to find book with ID: %s", bookID)
	}

	return nil
}

func (s *BookStorage) DeleteBook(bookID string) error {
	const query = `DELETE FROM books
	               WHERE id = $1`

	result, err := s.db.Exec(context.Background(), query, bookID)
	if err != nil {
		return fmt.Errorf("delete book: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed find book with ID: %s", bookID)
	}

	return nil
}

func (s *BookStorage) GetBook(bookID string) (*models.Book, error) {
	const query = `SELECT id, user_id, name, description, rating, status
	               FROM books WHERE id = $1`

	var book models.Book
	err := s.db.QueryRow(context.Background(), query, bookID).Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("failed to find book with ID: %s", bookID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &book, nil
}

func (s *BookStorage) GetUserBooks(userID string) ([]*models.Book, error) {
	const query = `SELECT id, user_id, name, description, rating, status FROM books
	               WHERE user_id = $1`

	rows, err := s.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users books: %w", err)
	}
	defer rows.Close()

	var userBooks []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
		if err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}

		userBooks = append(userBooks, &book)
	}

	return userBooks, nil
}
