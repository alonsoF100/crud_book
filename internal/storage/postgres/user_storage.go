package postgres

import (
	"context"
	"crud_book/internal/dto"
	"crud_book/internal/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *userStorage {
	return &userStorage{db: db}
}

func (s *userStorage) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {
	newID := uuid.New().String()

	const query = `INSERT INTO users (id, name, email) 
				   VALUES ($1, $2, $3)
				   RETURNING id, name, email`

	var user models.User
	err := s.db.QueryRow(ctx, query, newID, req.Name, req.Email).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (s *userStorage) GetUser(ctx context.Context, req dto.GetUserRequest) (*models.User, error) {
	const query = `SELECT id, name, email
	               FROM users WHERE id = $1`

	var user models.User
	err := s.db.QueryRow(ctx, query, req.UserID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("failed to find user with ID: %s", req.UserID)
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}

func (s *userStorage) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	const query = `SELECT id, name, email FROM users`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		var user models.User
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

func (s *userStorage) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error {
	const query = `DELETE FROM users WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.UserID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed find user with ID: %s", req.UserID)
	}

	return nil
}
