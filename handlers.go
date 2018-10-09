package main

// net/http:This is required for example if using constants such as http.StatusOK
// strconv: Implements conversions from and to String.
// strings: Package strings implements simple functions to manipulate UTF-8 encoded strings.
// Other Packages: https://golang.org/pkg/
import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Get All Fraud-Detection requests from Database.
func fetchFraudDetectionRequests(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	// Read from Database:
	// ?

	// gin.H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"message": "Fraud-Detection handler is not implemented yet",
	})
}

// Get a single Fraud-Detection from the Database.
func fetchFraudDetectionRequest(c *gin.Context) {
	// Read the input "id" param.
	paramID := c.Param("id")

	// Return a 404 response (NotFound).
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": strings.Join([]string{"Fraud-Detection not found, id:", paramID}, "")})
}

// Validate the Payment for possible Fraud
func validatePayment(c *gin.Context) {
	// Read an Integer param (from POST)
	// Atoi is used to convert string to int.
	intParam1, err := strconv.Atoi(c.PostForm("intParam1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "StatusBadRequest"})
		return
	}

	// Read a String param (from POST)
	strParam1 := c.PostForm("stParam1")
	if len(strParam1) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "StatusBadRequest"})
		return
	}

	// Insert to Database:
	// ?
	dummy := strings.Join([]string{strconv.Itoa(intParam1), strParam1}, ":")

	// Return a dummy created response.
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Fraud-Detection item created successfully!",
		"resourceId": strings.Join([]string{"Just Kidding, it is not not implemented yet.", dummy}, "")})
}

// Update the details of an Fraud-Detection.
func updateFraudDetectionRequests(c *gin.Context) {
	// Read input params just like in the validatePayment
	// ?

	// Update Fraud-Detection in Database:
	// ?

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Fraud-Detection handler is not implemented yet"})
}

// Delete a Fraud-Detection.
func deleteFraudDetectionRequests(c *gin.Context) {
	// Read input params just like in the validatePayment
	// ?

	// Delete Fraud-Detection in Database:
	// ?

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Fraud-Detection handler is not implemented yet"})
}
