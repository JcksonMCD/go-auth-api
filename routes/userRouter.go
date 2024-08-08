package routes

import (
	"github.com/JcksonMCD/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.getUsers())
	incomingRoutes.GET("/user:user_id", controller.getUserById())

}
