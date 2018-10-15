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

// Return All Fraud-Detection requests from Database.
func fetchFraudDetectionRequests(c *gin.Context) {
	// swagger:route GET /api/v1/payments/fraud-detection/ fetchFraudDetectionRequests
	//
	// Handler returning list of All Fraud-Detection requests from Database.
	//
	// List of All Fraud-Detection requests from Database
	//
	// Responses:
	//   200: repoResp
	//   403: forbidden

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
	// swagger:operation GET /api/v1/payments/fraud-detection/{id} fetchFraudDetectionRequest
	//
	// Returns a single Fraud-Detection from the Database.
	//
	// Could be info for any Fraud-Detection...
	//
	// ---
	// parameters:
	// - name: id
	//   in: path
	//   description: optional filter to obtain statistics by proper country code in ISO 3166-1 numeric
	//   type: integer
	//   required: true
	// responses:
	//   '200':
	//     description: "returns statistics about bought, only ordered, and returned products"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           title:
	//             description: title of product
	//             type: string
	//           bought:
	//             description: how many items of specific products were bought in the previous week
	//             type: integer
	//           ordered:
	//             description: how many items of specific products were only ordered in the previous week
	//             type: integer
	//           returned:
	//             description: how many items of specific products were returned in the previous week
	//             type: integer
	//         "required": ["bought", "returned", "ordered"]

	// Read the input "id" param.
	paramID := c.Param("id")

	// Return a 404 response (NotFound).
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": strings.Join([]string{"Fraud-Detection not found, id:", paramID}, "")})
}

// Validate the Payment for possible Fraud
func validatePayment(c *gin.Context) {
	// swagger:operation POST /api/v1/payments/fraud-detection/ validatePaymentRequest
	//
	// validatePayment: Validate the Payment for possible Fraud
	//
	// Could be info for any Fraud-Detection...
	//
	// ---
	// consumes:
	// - application/x-www-form-urlencoded
	// responses:
	//   '200':
	//     description: "returns statistics about bought, only ordered, and returned products"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           status:
	//             description: the respose status
	//             type: string
	//           message:
	//             description: the response message
	//             type: string
	//           resourceId:
	//             description: the id of the new
	//             type: string
	//         "required": ["status", "message"]

	// Read an Integer param (from POST)
	// Atoi is used to convert string to int.
	intParam1, err := strconv.Atoi(c.PostForm("intParam1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "StatusBadRequest"})
		return
	}

	// Read a String param (from POST)
	strParam1 := c.PostForm("strParam1")
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
	// swagger:operation PUT /api/v1/payments/fraud-detection/ updateFraudDetectionRequests
	//
	// updateFraudDetectionRequests: Update the details of an Fraud-Detection.
	//
	// Could be info for any Fraud-Detection...
	//
	// ---
	// parameters:
	// - name: resourceId
	//   in: query
	//   description: required parameter with desctiption
	//   type: integer
	// responses:
	//   '200':
	//     description: "returns statistics about bought, only ordered, and returned products"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           status:
	//             description: the respose status
	//             type: string
	//           message:
	//             description: the response message
	//             type: string
	//         "required": ["status", "message"]
	//   '404':
	//     description: "Not found"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           status:
	//             description: the respose status
	//             type: string
	//           message:
	//             description: the response message
	//             type: string
	//         "required": ["status", "message"]

	// Read input params just like in the validatePayment
	// ?

	// Update Fraud-Detection in Database:
	// ?

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Fraud-Detection handler is not implemented yet"})
}

// Delete a Fraud-Detection.
func deleteFraudDetectionRequests(c *gin.Context) {
	// swagger:operation DELETE /api/v1/payments/fraud-detection/ deleteFraudDetectionRequests
	//
	// deleteFraudDetectionRequests: Delete a Fraud-Detection.
	//
	// Could be info for any Fraud-Detection...
	//
	// ---
	// parameters:
	// - name: resourceId
	//   in: query
	//   description: required parameter with desctiption
	//   type: integer
	// responses:
	//   '200':
	//     description: "returns statistics about bought, only ordered, and returned products"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           status:
	//             description: the respose status
	//             type: string
	//           message:
	//             description: the response message
	//             type: string
	//         "required": ["status", "message"]
	//   '404':
	//     description: "Not found"
	//     schema:
	//       type: array
	//       items:
	//         type: object
	//         properties:
	//           status:
	//             description: the respose status
	//             type: string
	//           message:
	//             description: the response message
	//             type: string
	//         "required": ["status", "message"]

	// Read input params just like in the validatePayment
	// ?

	// Delete Fraud-Detection in Database:
	// ?

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Fraud-Detection handler is not implemented yet"})
}
