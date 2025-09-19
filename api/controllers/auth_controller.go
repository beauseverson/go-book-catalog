package controllers

import (
	"context"
	"time"
	"log"
	"errors"
	"net/http"

	"go-book-catalog/models"
	"go-book-catalog/utils"
	"go-book-catalog/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/golang-jwt/jwt/v5"
)

// get envar for JWT secret key using the utils.GetEnvVar function
var jwtKey = []byte(utils.GetEnvVar("JWT_SECRET"))

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JSON request body to get username and password
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Fetch user from the database
		collection := database.Client.Database("bookdb").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var storedUser models.User
		err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&storedUser)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			} else {
				log.Println("Error fetching user:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		// Check password
		if !utils.CheckPasswordHash(user.Password, storedUser.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": storedUser.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		// Sign the token with the secret key
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Println("Error signing token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Return the token
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}