package routes

import (
	controller "github.com/JcksonMCD/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func authRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
