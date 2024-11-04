package middleware

import (
	"dating-app-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function that checks for a valid JWT token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the user ID in the context for later use
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
