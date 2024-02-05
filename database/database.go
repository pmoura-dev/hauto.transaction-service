package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/pmoura-dev/hauto.transaction-service/config"
)

func GetConnection(config config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
