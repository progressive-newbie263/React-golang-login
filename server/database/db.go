package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

// Accept the connection string as a parameter
func Connect(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}