package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS account (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			nik VARCHAR(20) NOT NULL UNIQUE,
			no_hp VARCHAR(15) NOT NULL,
			no_rekening VARCHAR(20) NOT NULL UNIQUE,
			balance DECIMAL(15,2) NOT NULL DEFAULT 0
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
