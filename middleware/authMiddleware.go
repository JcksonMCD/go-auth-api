package middleware

import (
	"net/http"

	service "github.com/JcksonMCD/golang-jwt/service"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the token from the request header
		clientToken := c.Request.Header.Get("token")

		// check if the token is provided in the request header
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization header provided"})
			c.Abort() // Abort the request processing
			return
		}

		// validate the token
		claims, err := service.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort() // Abort the request processing
			return
		}

		// set context fields to those that have been returned
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("last_name", claims.Last_Name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_Type)
		c.Next()

	}
}
