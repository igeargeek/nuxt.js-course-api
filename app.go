package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"app/src/controllers"
	"app/src/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler        controllers.UserHandler
	MovieHandler       controllers.MovieHandler
	ReservationHandler controllers.ReservationHandler
}

func NewAppDatabase(userHandler controllers.UserHandler, movieHandler controllers.MovieHandler, reservationHandler controllers.ReservationHandler) App {
	return App{
		UserHandler:        userHandler,
		MovieHandler:       movieHandler,
		ReservationHandler: reservationHandler,
	}
}

func main() {
	mongoURI := fmt.Sprintf(os.Getenv("MONGO_URI"))
	app, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		log.Fatal("App initial error")
	}

	router := gin.Default()

	userRoute := router.Group("/users")
	routes.UserRoute(userRoute, app.UserHandler)

	movieRoute := router.Group("/movies")
	routes.MovieRoute(movieRoute, app.MovieHandler, app.ReservationHandler)

	reservationRoute := router.Group("/reservations")
	routes.ReservationRoute(reservationRoute, app.ReservationHandler)

	portListen := "8000"
	if port := os.Getenv("PORT"); port != "" {
		portListen = port
	}
	router.Run(":" + portListen)
}
