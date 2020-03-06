package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/icecreamhotz/movie-ticket/controllers"
	"github.com/icecreamhotz/movie-ticket/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB          *mongo.Client
	UserHandler controllers.UserHandler
}

func NewAppDatabase(db *mongo.Client, userHandler controllers.UserHandler) App {
	return App{
		DB:          db,
		UserHandler: userHandler,
	}
}

func main() {
	mongoURI := fmt.Sprintf("mongodb://mongodb:27017/")
	app, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		log.Fatal("App initial error")
	}

	databaseErr := app.DB.Ping(context.TODO(), nil)

	if databaseErr != nil {
		log.Fatal("databaseErr", databaseErr)
	}

	router := gin.Default()

	userRoute := router.Group("/user")
	routes.UserRoute(userRoute, app.UserHandler)

	router.Run(":" + os.Getenv("port"))
}
