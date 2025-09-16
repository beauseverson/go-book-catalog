package main

import (
	"go-book-catalog/database"
	"go-book-catalog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize default Gin Router
	router := gin.Default()

	// Setup the book routes
	routes.BookRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
