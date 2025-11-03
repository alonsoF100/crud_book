package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	db *pgx.Conn
}

func NewConnect() *Storage {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
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
	err := s.db.Close(context.Background())
	if err != nil {
		log.Fatalf("failed to disconnect %v", err)
	}
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("failed to find user with ID: %s", userID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}

func (s *Storage) GetAllUsers() ([]*User, error) {
	const query = `select id, name, email from users`

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*User
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

// переделать на базу

func (s *Storage) AddBook(userID, name, description string) (*Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book := &Book{
		ID:          uuid.New().String(),
		UserID:      userID,
		Name:        name,
		Description: description,
		Rating:      0,
		Status:      "want",
	}

	if _, exists := s.users[userID]; !exists {
		return nil, fmt.Errorf("пользователь с ID %s не найден", userID)
	}

	s.books[book.ID] = book

	return book, nil
}

func (s *Storage) UpdateBookStatus(bookID, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	validStatuses := map[string]bool{
		"want":     true,
		"reading":  true,
		"finished": true,
	}

	if !validStatuses[status] {
		return fmt.Errorf("неверный статус: %s. Допустимые значения: want, reading, finished", status)
	}

	book, exist := s.books[bookID]
	if !exist {
		return fmt.Errorf("неудалось найти книгу с ID: %s", bookID)
	}

	book.Status = status

	return nil
}

func (s *Storage) UpdateBookRating(bookID string, rating float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if rating < 0 || rating > 5 {
		return fmt.Errorf("рейтинг должен быть от 0 до 5")
	}

	book, exist := s.books[bookID]
	if !exist {
		return fmt.Errorf("неудалось найти книгу с ID: %s", bookID)
	}

	book.Rating = rating

	return nil
}

func (s *Storage) DeleteBook(bookID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.books[bookID]; !exist {
		return fmt.Errorf("неудалось найти книгу с ID: %s", bookID)
	}

	delete(s.books, bookID)

	return nil
}

func (s *Storage) GetBook(bookID string) (*Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, exist := s.books[bookID]
	if !exist {
		return nil, fmt.Errorf("неудалось найти книгу с ID: %s", bookID)
	}

	return book, nil
}

func (s *Storage) GetUserBooks(userID string) ([]*Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exist := s.users[userID]; !exist {
		return nil, fmt.Errorf("неудалось найти пользователя с ID: %s", userID)
	}

	var userBooks []*Book
	for _, book := range s.books {
		if book.UserID == userID {
			userBooks = append(userBooks, book)
		}
	}

	return userBooks, nil
}
