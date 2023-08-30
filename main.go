package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"myproject/kyc"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv" //// New import
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create KYC documents table if not exists
	err = createKYCDocumentTable(db)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Add CORS configuration using gorilla/handlers
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Adjust this to restrict origins
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	r.Use(corsHandler) // Use the CORS handler

	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Add KYC routes using the kyc package functions
	r.HandleFunc("/kyc/upload", func(w http.ResponseWriter, r *http.Request) {
		kyc.UploadKYCDocumentHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/kyc/status/{document_id}", func(w http.ResponseWriter, r *http.Request) {
		kyc.KYCStatusHandler(w, r, db)
	}).Methods("GET")

	// Add KYC routes using the kyc handler
	kycHandler := kyc.NewKYCHandler(db)

	r.HandleFunc("/kyc/upload", kycHandler.UploadKYCDocumentHandler).Methods("POST")
	r.HandleFunc("/kyc/status/{document_id}", kycHandler.KYCStatusHandler).Methods("GET")

	http.Handle("/", r)
	if err = http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  "Success",
		Message: "User created successfully!",
	}

	json.NewEncoder(w).Encode(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser User
	err = db.QueryRow("SELECT * FROM users WHERE username=$1", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	response := Response{
		Status:  "Success",
		Message: "Logged in successfully!",
	}

	json.NewEncoder(w).Encode(response)
}

func createKYCDocumentTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS kyc_documents (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			document_type VARCHAR(255) NOT NULL,
			document_data BYTEA NOT NULL,
			status VARCHAR(50) DEFAULT 'Pending'
		);
	`)
	return err
}
