package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"time"

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
type KYCUserStep3 struct {
	UniqueID       string `json:"unique_id"`
	AddressLine1   string `json:"address_line1"`
	AddressLine2   string `json:"address_line2"`
	City           string `json:"city"`
	StateProvince  string `json:"state_province"`
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
}
type KYCUserStep4 struct {
    UniqueID              string `json:"unique_id"`
    LivedAtAddress3Years  bool   `json:"lived_at_address_3_years"`  // true for Yes, false for No
}
type KYCUserStep5 struct {
    UniqueID   string `json:"unique_id"`
    IDType     string `json:"id_type"` // "Passport", "UK driving licence", or "EU national identity card"
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
// KYCHandlerStep3 handles the third step of the KYC process


	// Directly use the date in the ISO format "YYYY-MM-DD" provided by the frontend
	_, err = db.Exec("UPDATE kyc_users SET date_of_birth = $1 WHERE unique_id = $2", user.DateOfBirth, user.UniqueID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Date of birth saved successfully!"})

}
func KYCHandlerStep3(c *gin.Context) {
	var user KYCUserStep3
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE kyc_users SET address_line1 = $1, address_line2 = $2, city = $3, state_province = $4, postal_code = $5, country = $6 WHERE unique_id = $7", user.AddressLine1, user.AddressLine2, user.City, user.StateProvince, user.PostalCode, user.Country, user.UniqueID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Address information saved successfully!"})
}
// KYCHandlerStep4 handles the fourth step of the KYC process
func KYCHandlerStep4(c *gin.Context) {
    var user KYCUserStep4
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Storing the user's answer into the database.
    _, err = db.Exec("UPDATE kyc_users SET lived_at_address_3_years = $1 WHERE unique_id = $2", user.LivedAtAddress3Years, user.UniqueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Address duration information saved successfully!"})
}
func KYCHandlerStep5(c *gin.Context) {
    var user KYCUserStep5
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate ID type
    validIDTypes := []string{"Passport", "UK driving licence", "EU national identity card"}
    isValidType := false
    for _, validType := range validIDTypes {
        if user.IDType == validType {
            isValidType = true
            break
        }
    }
    if !isValidType {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID type selected."})
        return
    }

    // Store the ID type in the database
    _, err = db.Exec("UPDATE kyc_users SET id_type = $1 WHERE unique_id = $2", user.IDType, user.UniqueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ID type saved successfully!"})
}

var s3Session = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-west-2"),  // e.g., "us-west-1"
	// Other AWS session configuration may be added here
}))

// GetPresignedURLHandler generates a presigned URL for S3 uploads
func GetPresignedURLHandler(c *gin.Context) {
	var request struct {
		UniqueID string `json:"unique_id"`
		IDType   string `json:"id_type"`
		FileName string `json:"file_name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s3svc := s3.New(s3Session)
	req, _ := s3svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("kycsolution"),
		Key:    aws.String(fmt.Sprintf("%s/%s/%s", request.UniqueID, request.IDType, request.FileName)),
	})
	presignedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate presigned URL."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"presigned_url": presignedURL})
}