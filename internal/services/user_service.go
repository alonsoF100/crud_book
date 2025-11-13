package services

import (
	"crud_book/internal/dto"
	"crud_book/internal/models"
	"crud_book/internal/storage"
)

type userServiceImpl struct {
	repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *userServiceImpl {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(req dto.CreateUserRequest) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateUser(req)
}

func (s *userServiceImpl) GetUser(req dto.GetUserRequest) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUser(req)
}

func (s *userServiceImpl) DeleteUser(req dto.DeleteUserRequest) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteUser(req)
}

func (s *userServiceImpl) GetAllUsers() ([]*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetAllUsers()
}
