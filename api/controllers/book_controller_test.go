package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"go-book-catalog/database"
	"go-book-catalog/testutils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	mongoContainer, uri, err := testutils.SetupTestMongoDB()
	if err != nil {
		os.Exit(1)
	}
	database.ConnectDB(uri)

	exitCode := m.Run()

	if err := mongoContainer.Terminate(context.Background()); err != nil {
		os.Exit(1)
	}
	os.Exit(exitCode)
}

func TestCreateBook(t *testing.T) {
	// 1. Create a Gin test environment
	router := gin.Default()
	router.POST("/books", CreateBook())

	// 2. Create a mock book payload
	mockBook := gin.H{
		"title":  "Test Book",
		"author": "Test Author",
		"genre":  "Test Genre",
		"year":   2024,
	}
	jsonValue, _ := json.Marshal(mockBook)

	// 3. Create a simulated HTTP request
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// 4. Create a recorder to capture the response
	w := httptest.NewRecorder()

	// 5. Serve the request to the router
	router.ServeHTTP(w, req)

	// 6. Assert the results
	assert.Equal(t, http.StatusCreated, w.Code)

	// Optional: Parse the response body to check for the ID
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["message"], "created successfully")
}
