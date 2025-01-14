package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/go-libsql"
)

func NewConnection() (*sqlx.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
