package apis

import (
	"database/sql"
	"log"
)

var db *sql.DB

// InitDB initializes the database connection for the KYC API
func InitDB(database *sql.DB) {
	db = database

	// Create table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS kyc_users (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		middle_name TEXT,
		last_name TEXT NOT NULL
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating kyc_users table: %v", err)
	}
}
