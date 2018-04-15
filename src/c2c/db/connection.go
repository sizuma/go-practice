package db

import (
	"c2c/config"
	"database/sql"
)

// Connect return new database connection
func Connect() (*sql.DB, error) {
	projectConfig := config.Get()
	return sql.Open("mysql", projectConfig.Database.URL())
}
