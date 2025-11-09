package services

import (
	"crud_book/internal/models"
	"crud_book/internal/storage/postgres"
)

type UserService struct {
	storage *postgres.Storage
}

func NewUserService(storage *postgres.Storage) *UserService {
	return &UserService{storage: storage}
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.CreateUser(name, email)
}

func (s *UserService) GetUser(userID string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.GetUser(userID)
}

func (s *UserService) DeleteUser(userID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.DeleteUser(userID)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.storage.GetAllUsers()
}
