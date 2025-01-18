package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/CyberBoyzzz/Framezzz/controllers"
	"github.com/CyberBoyzzz/Framezzz/middlewares" 
)

// SetupRoutes registers all routes for the application
func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
