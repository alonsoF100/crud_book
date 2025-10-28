package storage

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Book struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Status      string  `json:"status"` // будет три статуса "want" "reading" "finished"
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Storage struct {
	users map[string]*User
	books map[string]*Book
	mu    sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]*User),
		books: make(map[string]*Book),
		mu:    sync.RWMutex{},
	}
}

func (s *Storage) AddUser(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}

	for _, existingUser := range s.users {
		if existingUser.Email == email {
			return nil, fmt.Errorf("пользователь с email %s уже существует", email)
		}
	}

	s.users[user.ID] = user

	return user, nil
}

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

func (s *Storage) DeleteUser(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.users[userID]; !exist {
		return fmt.Errorf("неудалось найти пользователя с ID: %s", userID)
	}

	for bookID, book := range s.books {
		if book.UserID == userID {
			delete(s.books, bookID)
		}
	}

	delete(s.users, userID)

	return nil
}

func (s *Storage) GetUser(userID string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exist := s.users[userID]
	if !exist {
		return nil, fmt.Errorf("неудалось найти пользователя с ID: %s", userID)
	}

	return user, nil
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

func (s *Storage) GetAllUsers() []*User {
	var users []*User
	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}
