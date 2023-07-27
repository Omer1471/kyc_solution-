package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Task represents a single to-do task
type Task struct {
	ID    int
	Title string
}

// User represents a user for registration and authentication
type User struct {
	ID       int
	Username string
	Password []byte
}

var db *sql.DB

func main() {
	// Database connection settings
	connStr := "postgres://postgres:Rockerz88@localhost:5432/postgres?sslmode=disable"

	// Connect to the database
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create the "tasks" table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "users" table for user registration
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password BYTEA NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/register", registerHandler) // New route for user registration
	http.HandleFunc("/login", loginHandler)       // New code for user login

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	log.Println("Server started on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all tasks from the database
	// ... (Your existing code for fetching tasks)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// ... (Your existing code for adding tasks)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// ... (Your existing code for deleting tasks)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the username and password from the form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Insert the user data into the "users" table
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to insert user data", http.StatusInternalServerError)
		return
	}

	// Redirect the user to the home page after successful registration
	http.Redirect(w, r, "/", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the username and password from the form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Fetch the user with the given username from the database
	var dbPassword []byte
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&dbPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare the stored password hash with the provided password
	err = bcrypt.CompareHashAndPassword(dbPassword, []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Successful login, redirect the user to the home page
	http.Redirect(w, r, "/", http.StatusFound)
}

