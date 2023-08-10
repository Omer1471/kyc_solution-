package kyc

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

// Response represents the structure of an API response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// KYCDocument represents the structure of a KYC document
type KYCDocument struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	DocumentType string `json:"document_type"`
	Status       string `json:"status"`
}

// UploadKYCDocumentHandler handles document uploads
func UploadKYCDocumentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse the incoming multipart/form-data request
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the uploaded file from the form data
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Extract other form values if needed
	userID := r.FormValue("user_id")
	documentType := r.FormValue("document_type")

	// Read the file data
	fileData := make([]byte, fileHeader.Size)
	_, err = file.Read(fileData)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Insert document details into the database
	_, err = db.Exec(
		"INSERT INTO kyc_documents (user_id, document_type, document_data) VALUES ($1, $2, $3)",
		userID, documentType, fileData,
	)
	if err != nil {
		http.Error(w, "Error inserting document", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	response := Response{
		Status:  "Success",
		Message: "Document uploaded successfully!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// KYCStatusHandler handles KYC document status requests
func KYCStatusHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Implement document status retrieval logic and database interaction
	// ...

	// Respond with document status
	// ...
    



}
