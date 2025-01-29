// 1. infrastructure/database/connection.go
package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectPostgres(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
