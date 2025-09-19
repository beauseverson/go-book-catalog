package utils

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// GetEnvVar loads .env and returns the value for the given key
func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// CheckPasswordHash compares a plaintext password with a hashed password and returns true if they match
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}