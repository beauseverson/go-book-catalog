package controllers

import (
	"context"
	"go-book-catalog/database"
	"go-book-catalog/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getBookCollection() *mongo.Collection {
	return database.Client.Database("bookdb").Collection("books")
}

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := getBookCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var book models.Book
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := bookCollection.InsertOne(ctx, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new book"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book_id": result.InsertedID})
	}
}

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := getBookCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var books []models.Book
		defer cancel()

		cursor, err := bookCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching books"})
			return
		}

		if err = cursor.All(ctx, &books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding books"})
			return
		}

		// check if books is empty and return an empty array instead of null
		if len(books) == 0 {
			c.JSON(http.StatusOK, []models.Book{})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := getBookCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookId")
		var book models.Book
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(bookID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		err = bookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching the book"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := getBookCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookId")
		var updatedBook models.Book
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(bookID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := bson.M{
			"$set": updatedBook,
		}

		result, err := bookCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating the book"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	}
}

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := getBookCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookId")
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(bookID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting the book"})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}
