package routes

import (
	controller "go-lang-jwt-project/controllers"
	"go-lang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func USerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
