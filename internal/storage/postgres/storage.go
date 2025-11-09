package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

type Book struct {
	ID          string  `json:"id" db:"id"`
	UserID      string  `json:"user_id" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Rating      float64 `json:"rating" db:"rating"`
	Status      string  `json:"status" db:"status"`
}

type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type Storage struct {
	db *pgxpool.Pool
}

func NewConnect() *Storage {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("LOCAL_DATABASE_URL"))
	if err != nil {
		log.Fatalf("cant connect to database %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("failed ping %v", err)
	}

	return &Storage{db: db}
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) RunMigrations() error {
	godotenv.Load()

	db, err := goose.OpenDBWithDriver("pgx", os.Getenv("LOCAL_DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("open db for migrations: %w", err)
	}
	defer db.Close()

	err = goose.Up(db, "internal/storage/postgres/migrations")
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

func (s *Storage) AddUser(name, email string) (*User, error) {
	newID := uuid.New().String()

	const query = `INSERT INTO users (id, name, email) 
				   VALUES ($1, $2, $3)
				   RETURNING id, name, email`

	var user User
	err := s.db.QueryRow(context.Background(), query, newID, name, email).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (s *Storage) GetUser(userID string) (*User, error) {
	const query = `SELECT id, name, email
	               FROM users WHERE id = $1`

	var user User
	err := s.db.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("failed to find user with ID: %s", userID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}

func (s *Storage) GetAllUsers() ([]*User, error) {
	const query = `SELECT id, name, email FROM users`

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

func (s *Storage) DeleteUser(userID string) error {
	const query = `DELETE FROM users WHERE id = $1`

	result, err := s.db.Exec(context.Background(), query, userID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed find user with ID: %s", userID)
	}

	return nil
}

func (s *Storage) AddBook(userID, name, description string) (*Book, error) {
	newID := uuid.New().String()

	const query = `INSERT INTO books (id, user_id, name, description)
	               VALUES ($1, $2, $3, $4)
				   RETURNING id, user_id, name, description, rating, status`
	var book Book
	err := s.db.QueryRow(context.Background(), query, newID, userID, name, description).
		Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
	if err != nil {
		return nil, fmt.Errorf("create book: %w", err)
	}

	return &book, nil
}

func (s *Storage) UpdateBookStatus(bookID, status string) error {
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

func (s *Storage) UpdateBookRating(bookID string, rating float64) error {
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

func (s *Storage) DeleteBook(bookID string) error {
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

func (s *Storage) GetBook(bookID string) (*Book, error) {
	const query = `SELECT id, user_id, name, description, rating, status
	               FROM books WHERE id = $1`

	var book Book
	err := s.db.QueryRow(context.Background(), query, bookID).Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("failed to find book with ID: %s", bookID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &book, nil
}

func (s *Storage) GetUserBooks(userID string) ([]*Book, error) {
	const query = `SELECT id, user_id, name, description, rating, status FROM books
	               WHERE user_id = $1`

	rows, err := s.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users books: %w", err)
	}
	defer rows.Close()

	var userBooks []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Name, &book.Description, &book.Rating, &book.Status)
		if err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}

		userBooks = append(userBooks, &book)
	}

	return userBooks, nil
}
