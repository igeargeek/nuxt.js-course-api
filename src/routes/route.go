package routes

import (
	"app/src/controllers"

	"github.com/gin-gonic/gin"
)

// var userHandler controllers.UserHandler

func UserRoute(router *gin.RouterGroup, userHandler controllers.UserHandler) {
	router.POST("/register", userHandler.RegisterUserPost)
	router.GET("/", userHandler.ListOfUsersGet)
}
