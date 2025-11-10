package services

import (
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type UserService struct {
	repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateUser(name, email)
}

func (s *UserService) GetUser(userID string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUser(userID)
}

func (s *UserService) DeleteUser(userID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteUser(userID)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetAllUsers()
}
