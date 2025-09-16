package main

import (
	"go-book-catalog/database"
	"go-book-catalog/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize default Gin Router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Your React app's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	}))

	// Setup the book routes
	routes.BookRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
