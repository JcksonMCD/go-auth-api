package middleware

import (
	"log"
	"net/http"

	service "github.com/JcksonMCD/golang-jwt/service"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the request header
		clientToken := c.Request.Header.Get("token")

		// Check if the token is provided in the request header
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization header provided"})
			c.Abort() // Abort the request processing
			return
		}

		// Validate the token
		claims, err := service.ValidateToken(clientToken)
		if err != "" {
			log.Printf("Token validation error: %s", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort() // Abort the request processing
			return
		}

		// Set context fields to those that have been returned
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("last_name", claims.Last_Name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_Type)
		c.Next()
	}
}
