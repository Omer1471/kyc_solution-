package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type KYCUser struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

// KYCHandler handles KYC submissions and stores them in the database
func KYCHandler(c *gin.Context) {
	var kycUser KYCUser
	err := c.ShouldBindJSON(&kycUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if kycUser.FirstName == "" || kycUser.MiddleName == "" || kycUser.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All name fields are required."})
		return
	}

	uniqueID := generateRandomID(16) // This will generate a 32-character long unique ID

	_, err = db.Exec("INSERT INTO kyc_users (unique_id, first_name, middle_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id", uniqueID, kycUser.FirstName, kycUser.MiddleName, kycUser.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User information saved successfully!"})
}
