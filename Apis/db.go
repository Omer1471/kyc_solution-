package apis

import (
	"database/sql"
	"log"
	"crypto/rand"
	"encoding/hex"
)

var db *sql.DB

// InitDB initializes the database connection for the KYC API
func InitDB(database *sql.DB) {
	db = database

	// Drop the kyc_users table
	dropTable := `
	DROP TABLE IF EXISTS kyc_users;
	`
	_, err := db.Exec(dropTable)
	if err != nil {
		log.Fatalf("Error dropping kyc_users table: %v", err)
	}

	// Recreate the table with the unique_id column
	createTable := `
	CREATE TABLE kyc_users (
		id SERIAL PRIMARY KEY,
		unique_id VARCHAR(255) UNIQUE NOT NULL,
		first_name TEXT NOT NULL,
		middle_name TEXT,
		last_name TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating kyc_users table: %v", err)
	}
}

// generateRandomID generates a cryptographically secure random ID of given length
func generateRandomID(length int) string {
	bytes := make([]byte, length/2)  // Divide by 2 because each byte will be represented by 2 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		// Handle error here
		log.Fatalf("Error generating random ID: %v", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}
