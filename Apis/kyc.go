package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type KYCUserStep1 struct {
	ID         int    `json:"id"`
	UniqueID   string `json:"unique_id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type KYCUserStep2 struct {
	UniqueID     string `json:"unique_id"`
	DateOfBirth  string `json:"date_of_birth"`
}

// KYCHandlerStep1 handles the first step of the KYC process
func KYCHandlerStep1(c *gin.Context) {
	var user KYCUserStep1
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uniqueID := generateRandomID(16)
	_, err = db.Exec("INSERT INTO kyc_users (unique_id, first_name, middle_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id", uniqueID, user.FirstName, user.MiddleName, user.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User information saved successfully!", "unique_id": uniqueID})
}

// KYCHandlerStep2 handles the second step of the KYC process
func KYCHandlerStep2(c *gin.Context) {
	var user KYCUserStep2
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Directly use the date in the ISO format "YYYY-MM-DD" provided by the frontend
	_, err = db.Exec("UPDATE kyc_users SET date_of_birth = $1 WHERE unique_id = $2", user.DateOfBirth, user.UniqueID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Date of birth saved successfully!"})
}
