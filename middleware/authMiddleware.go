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
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort() // Abort the request processing
			return
		}
	}
}
