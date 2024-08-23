package routes

import (
	controller "github.com/JcksonMCD/golang-jwt/controllers"
	middleware "github.com/JcksonMCD/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUserById())

}
