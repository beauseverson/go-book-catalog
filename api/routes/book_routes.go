package routes

import (
	"go-book-catalog/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	// new change
	bookRoutes := router.Group("/books")
	bookRoutes.POST("", controllers.CreateBook())
	bookRoutes.GET("", controllers.GetBooks())
	bookRoutes.GET("/:bookId", controllers.GetBook())
	bookRoutes.PUT("/:bookId", controllers.UpdateBook())
	bookRoutes.DELETE("/:bookId", controllers.DeleteBook())
}
