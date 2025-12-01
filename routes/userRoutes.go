package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func UserRouter(r *gin.Engine) {

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.UserLoginHandler)
	r.GET("/logout", controllers.UserLogoutHandler)
}