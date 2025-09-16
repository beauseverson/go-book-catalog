package main

import (
	"go-book-catalog/database"
	"go-book-catalog/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	}))

	// Setup the book routes
	routes.BookRoutes(router)

	return router
}

func getEnvVar(key string) string {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Return the value of the environment variable or an empty string if not found
	return os.Getenv(key)
}

func main() {
	// Connect to the database
	database.ConnectDB(getEnvVar("MONGODB_URI"))

	// Initialize default Gin Router
	router := setupRouter()

	// Start the server on port 8080
	router.Run(":8080")
}
