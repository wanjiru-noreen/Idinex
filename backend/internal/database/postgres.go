package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Connect creates a connection to the PostgreSQL database.
func Connect(databaseURL string) (*sql.DB, error) {
	// Open a connection using the PostgreSQL driver.
	db, err := sql.Open("postgres", databaseURL)

	// Return an error if the connection cannot be opened.
	if err != nil {
		return nil, err
	}

	// Verify that the database is reachable.
	err = db.Ping()

	// Return an error if the database cannot be reached.
	if err != nil {
		return nil, err
	}

	// Return the database connection.
	return db, nil
}
