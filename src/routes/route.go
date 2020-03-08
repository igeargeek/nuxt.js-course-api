package routes

import (
	"app/src/controllers"

	"github.com/gin-gonic/gin"
)

// var userHandler controllers.UserHandler

func UserRoute(router *gin.RouterGroup, userHandler controllers.UserHandler) {
	router.POST("/register", userHandler.RegisterUserPost)
	router.POST("/login", userHandler.LoginUserPost)
	router.POST("/refresh-token", userHandler.RefreshTokenPost)
	router.GET("/payload", userHandler.PayloadTokenGet)
}

func MovieRoute(router *gin.RouterGroup, movieHandler controllers.MovieHandler, reservationHandler controllers.ReservationHandler) {
	router.POST("", movieHandler.CreateMoviePost)
	router.GET("/:id", movieHandler.ShowOneMovieGet)
	router.DELETE("/:id", movieHandler.RemoveOneMovieDelete)
	router.PUT("/:id", movieHandler.EditMoviePut)
	router.GET("", movieHandler.ShowAllMovieGet)
	router.POST("/:id/_reseve", reservationHandler.CreateReservationPost)
}

func ReservationRoute(router *gin.RouterGroup, reservationHandler controllers.ReservationHandler) {
	router.GET("/:id", reservationHandler.ShowOneReservationGet)
	router.GET("", reservationHandler.ShowAllReservationGet)
}
