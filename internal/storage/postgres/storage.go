package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewConnect() *Storage {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = os.Getenv("LOCAL_DATABASE_URL")
	}
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("cant connect to database %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("failed ping %v", err)
	}

	return &Storage{db: db}
}

func (s *Storage) DB() *pgxpool.Pool {
	return s.db
}

func (s *Storage) Close() {
	s.db.Close()
}
