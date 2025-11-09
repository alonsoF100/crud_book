package services

import "crud_book/internal/storage/postgres"

type UserService struct {
	storage *postgres.Storage
}

func NewUserService(storage *postgres.Storage) *UserService {
	return &UserService{storage: storage}
}
