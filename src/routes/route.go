package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/controllers"
)

// var userHandler controllers.UserHandler

func UserRoute(router *gin.RouterGroup, userHandler controllers.UserHandler) {
	router.POST("/register", userHandler.RegisterUserPost)
}
