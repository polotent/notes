package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	DATABASE_HOST     = "localhost"
	DATABASE_PORT     = 5432
	DATABASE_USER     = "jkr"
	DATABASE_PASSWORD = "jkr"
	DATABASE_NAME     = "notes"
)

func init() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DATABASE_HOST, DATABASE_PORT, DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := DB.Close(); err == nil {
			log.Fatalf("Database connection close error: %s", err)
		}
	}()
}
