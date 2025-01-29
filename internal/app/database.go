package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

// ConnectDatabase подключается к базе данных PostgreSQL
func ConnectDatabase(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
