package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewConnection() (*sqlx.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("libsql", dbUrl)
	if err != nil {
		log.Println("Error connecting to database", err)
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
