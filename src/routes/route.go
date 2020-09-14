package routes

import (
	"app/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/src/middlewares"
)

// var userHandler controllers.UserHandler

func UserRoute(router *gin.RouterGroup, userHandler controllers.UserHandler) {
	router.POST("/register", userHandler.RegisterUserPost)
	router.POST("/login", userHandler.LoginUserPost)
	router.POST("/refresh-token", userHandler.RefreshTokenPost)
	router.GET("/", userHandler.ShowAllUserGet)
	router.DELETE("/", userHandler.DeleteAllDelete)
	router.GET("/payload", middlewares.AuthMiddleware(userHandler.Service), userHandler.PayloadTokenGet)
}

func MovieRoute(router *gin.RouterGroup, userHandler controllers.UserHandler, movieHandler controllers.MovieHandler, reservationHandler controllers.ReservationHandler) {
	router.POST("", movieHandler.CreateMoviePost)
	router.GET("/:id", movieHandler.ShowOneMovieGet)
	router.DELETE("/:id", movieHandler.RemoveOneMovieDelete)
	router.PUT("/:id", movieHandler.EditMoviePut)
	router.GET("", movieHandler.ShowAllMovieGet)
	router.POST("/:id/_reseve", middlewares.AuthMiddleware(userHandler.Service), reservationHandler.CreateReservationPost)
}

func ReservationRoute(router *gin.RouterGroup, userHandler controllers.UserHandler, reservationHandler controllers.ReservationHandler) {
	router.GET("/:id", middlewares.AuthMiddleware(userHandler.Service), reservationHandler.ShowOneReservationGet)
	router.GET("", middlewares.AuthMiddleware(userHandler.Service), reservationHandler.ShowAllReservationGet)
}
