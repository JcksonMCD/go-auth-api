package service

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// MatchUserTypeToID checks if the user type matches the user ID and verifies authorization.
func MatchUserTypeToID(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	// If user is a regular user and uid does not match the given userId, return an authorization error.
	if userType == "USER" && uid != userId {
		return errors.New("Unauthorized to access this resource")
	}

	return CheckUserType(c, userType)
}

// CheckUserType verifies if the user has the required role to access the resource (ADMIN).
func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")

	if userType != role {
		return errors.New("Unauthorized to access this resource")
	}

	return nil
}
