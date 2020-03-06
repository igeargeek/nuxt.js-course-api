package routes

import (
	"app/src/controllers"

	"github.com/gin-gonic/gin"
)

// var userHandler controllers.UserHandler

func UserRoute(router *gin.RouterGroup, userHandler controllers.UserHandler) {
	router.POST("/register", userHandler.RegisterUserPost)
}

func MovieRoute(router *gin.RouterGroup, movieHandler controllers.MovieHandler) {
	router.POST("/", movieHandler.CreateMoviePost)
}
