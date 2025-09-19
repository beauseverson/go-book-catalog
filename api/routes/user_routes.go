package routes

import (
	"go-book-catalog/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// new change
	authRoutes := router.Group("/auth")
	authRoutes.POST("/login", controllers.UserLogin())
	// authRoutes.POST("/register", controllers.RegisterHandler())
}