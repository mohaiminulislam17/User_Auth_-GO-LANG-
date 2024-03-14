package routes

import(

	controller "go-lang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/Signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}