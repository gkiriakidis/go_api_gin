// go-api-gin example.
//
// The purpose of this application is to provide examples
// use cases describing how to create a CRUD API, use unit testing,
// create automated documentation, and much, much more... stay tuned...
//
//     Schemes: http, https
//     Host: 127.0.0.1:8080
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default one shipped with Gin.
	router := gin.Default()

	// Set Custom Headers by creating a custom Middleware
	router.Use(setCustomHeaders())

	// Serve the frontend (static files)
	router.Static("/docs", "./docs")

	// Setup route group for the API (v1).
	v1 := router.Group("/api/v1/payments/fraud-detection")
	{
		v1.POST("/", validatePayment)
		v1.GET("/", fetchFraudDetectionRequests)
		v1.GET("/:id", fetchFraudDetectionRequest)
		v1.PUT("/:id", updateFraudDetectionRequests)
		v1.DELETE("/:id", deleteFraudDetectionRequests)
	}
	router.Run()
}

func setCustomHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	}
}
