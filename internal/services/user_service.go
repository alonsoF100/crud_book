package services

import (
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type userServiceImpl struct {
	repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *userServiceImpl {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(name, email string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateUser(name, email)
}

func (s *userServiceImpl) GetUser(userID string) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUser(userID)
}

func (s *userServiceImpl) DeleteUser(userID string) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteUser(userID)
}

func (s *userServiceImpl) GetAllUsers() ([]*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetAllUsers()
}
