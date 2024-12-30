/*
Database related services
*/

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./simple.db") // Path to SQLite file. Normally exported as env variable.
	if err != nil {
		return err
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
