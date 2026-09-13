package middleware

import (
	"net/http"

	"employee-management-backend/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtValidator *auth.JWTValidator) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("ACCESS_TOKEN")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		claims, err := jwtValidator.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.Subject)
		c.Set("username", claims.Username)
		c.Set("client_id", claims.ClientID)

		c.Next()
	}
}
