/*
 Every Go program starts with a package declaration.
 Packages are used to organize related go source code files into a single unit and make them reusable.
 The package "main" is a special go package that is used with programs that are meant to be executable.
 There are two types of programs in Go - Executable Programs and Libraries.
*/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default one shipped with Gin.
	router := gin.Default()

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
