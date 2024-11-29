package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pawel-mazurkiewicz/credit-card-validator/handlers"
)

func main() {
	router := gin.Default()

	// Middleware (optional)
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	router.POST("/validate", handlers.ValidateCardHandler)

	// Start the server
	router.Run(":8080") // Default port is 8080
}
