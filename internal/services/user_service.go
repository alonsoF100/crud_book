package services

import (
	"context"
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

func (s *userServiceImpl) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.CreateUser(ctx, req)
}

func (s *userServiceImpl) GetUser(ctx context.Context, req dto.GetUserRequest) (*models.User, error) {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.GetUser(ctx, req)
}

func (s *userServiceImpl) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error {
	// TODO Добавить валидацию и бизнес логику
	return s.repo.DeleteUser(ctx, req)
}

func (s *userServiceImpl) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	//TODO Добавить валидацию и бизнес логику
	return s.repo.GetAllUsers(ctx)
}
