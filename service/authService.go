package service

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToID(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err := nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorised to access this resource")
		return err
	}
	err = CheckUserType(c, userType)

	return err
}
