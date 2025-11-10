package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

type Migrator struct {
	migrationPath string
}

func NewMigrator(migrationPath string) *Migrator {
	return &Migrator{migrationPath: migrationPath}
}

func (m *Migrator) RunMigrations() error {
	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = os.Getenv("LOCAL_DATABASE_URL")
	}
	db, err := goose.OpenDBWithDriver("pgx", databaseURL)
	if err != nil {
		return fmt.Errorf("open db for migrations: %w", err)
	}
	defer db.Close()

	err = goose.Up(db, m.migrationPath)
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
