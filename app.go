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
	UserHandler  controllers.UserHandler
	MovieHandler controllers.MovieHandler
}

func NewAppDatabase(userHandler controllers.UserHandler, movieHandler controllers.MovieHandler) App {
	return App{
		UserHandler:  userHandler,
		MovieHandler: movieHandler,
	}
}

func main() {
	mongoURI := fmt.Sprintf(os.Getenv("mongo_uri"))
	app, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		log.Fatal("App initial error")
	}

	router := gin.Default()

	userRoute := router.Group("/users")
	routes.UserRoute(userRoute, app.UserHandler)

	movieRoute := router.Group("/movies")
	routes.MovieRoute(movieRoute, app.MovieHandler)

	router.Run(":" + os.Getenv("port"))
}
