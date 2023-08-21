package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/polotent/notes/src/backend/internal/app/notes/config"
)

const (
	POSTGRES = "postgres"
)

func ConnectDatabase(config *config.Config) (*sql.DB, error) {
	var connString string
	if config.Database.Dialect == POSTGRES {
		connString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBname,
		)
	} else {
		return nil, fmt.Errorf("Unsupported database dialect provided in config: %s", config.Database.Dialect)
	}

	var err error
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("Database open connection error: %s", err)
	}

	return db, nil
}
