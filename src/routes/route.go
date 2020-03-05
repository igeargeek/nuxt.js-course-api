package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/controllers"
)

var serviceHandler controllers.ServiceHandler

func UserRoute(router *gin.RouterGroup) {
	router.POST("/register", serviceHandler.RegisterUserPost)
}
