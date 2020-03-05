package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/controllers"
)

func UserRoute(router *gin.RouterGroup, serviceHandler *controllers.ServiceHandler) {
	router.POST("/register", serviceHandler.RegisterUserPost)
}
