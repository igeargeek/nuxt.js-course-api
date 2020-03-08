package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"app/src/controllers"
	"app/src/middlewares"
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

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	mongoURI := fmt.Sprintf(os.Getenv("MONGO_URI"))
	app, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		log.Fatal("App initial error")
	}

	router := gin.Default()
	router.Use(CORS)

	userRoute := router.Group("/users")
	routes.UserRoute(userRoute, app.UserHandler)

	movieRoute := router.Group("/movies")
	routes.MovieRoute(movieRoute, app.MovieHandler, app.ReservationHandler)

	reservationRoute := router.Group("/reservations", middlewares.AuthMiddleware(app.UserHandler.Service))
	routes.ReservationRoute(reservationRoute, app.ReservationHandler)

	portListen := "8000"
	if port := os.Getenv("PORT"); port != "" {
		portListen = port
	}
	router.Run(":" + portListen)
}
