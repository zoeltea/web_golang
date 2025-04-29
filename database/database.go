package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// postgresql://postgres:caEMpPeMZlEpOwMsGZzIbXouhpTfKGpW@turntable.proxy.rlwy.net:58393/railway
// PGPASSWORD=caEMpPeMZlEpOwMsGZzIbXouhpTfKGpW psql -h turntable.proxy.rlwy.net -U postgres -p 58393 -d railway

func InitDB() (*sql.DB, error) {

	fmt.Println("host: ", os.Getenv("PGHOST"))
	fmt.Println("port: ", os.Getenv("PGPORT"))
	fmt.Println("user: ", os.Getenv("PGUSER"))
	fmt.Println("password: ", os.Getenv("PGPASSWORD"))
	fmt.Println("host: databasename", os.Getenv("PGDATABASE"))
	connStr := fmt.Sprintf(
		// "host=${{ Postgres.DATABASE_URL }} port=5432 user=postgres password=caEMpPeMZlEpOwMsGZzIbXouhpTfKGpW dbname=railway sslmode=disable",
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
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
