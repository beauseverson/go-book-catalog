package main

import (
	"context"
	"go-book-catalog/database"
	"go-book-catalog/testutils"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

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

func TestBookCatalogAPI(t *testing.T) {
	// 1. Set up the test server
	router := setupRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	// 2. Make an HTTP request to the test server's URL
	res, err := http.Get(ts.URL + "/books")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	// 3. Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// 4. Assert the results
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, strings.Contains(string(body), "[]"), "Expected an empty book list initially")
}
