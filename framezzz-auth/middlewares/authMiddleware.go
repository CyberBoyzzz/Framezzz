package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/CyberBoyzzz/Framezzz/utils" 
)

// AuthMiddleware protects routes by validating the JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Attach claims to context for further use
		c.Set("userID", claims["userID"])
		c.Next()
	}
}
