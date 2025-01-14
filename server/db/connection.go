package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() (*sqlx.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("sqlite3", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
