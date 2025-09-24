package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDB(connectionString string) (*Database, error) {

	// Open DB connection
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// configure connection pool
	db.SetMaxOpenConns(25)                 // Limit max simultaneous connection
	db.SetMaxIdleConns(5)                  // Keep some connection ready
	db.SetConnMaxLifetime(5 * time.Minute) // Refresh conn periodically

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
