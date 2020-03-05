package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/controllers"
	"github.com/icecreamhotz/movie-ticket/routes"
)

type App struct {
	Handler *controllers.ServiceHandler
}

func NewApplication(handler *controllers.ServiceHandler) App {
	return App{
		Handler: handler,
	}
}

func main() {
	// mongoURI := fmt.Sprintf("mongodb+srv://%s:<%s>@cluster0-zz20n.mongodb.net/test?retryWrites=true&w=majority", "icecreamhotz", "zazaza1b")
	mongoURI := fmt.Sprintf("mongodb://mongodb:27017/")
	app, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		log.Fatal("failed to create event: %s\n", err)
	}

	e := app.Handler.DB.Ping(context.TODO(), nil)

	if e != nil {
		log.Fatal(e)
	}

	router := gin.Default()

	userRoute := router.Group("/user")
	routes.UserRoute(userRoute, app.Handler)

	router.Run(":8000")
}
