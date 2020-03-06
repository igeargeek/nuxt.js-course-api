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
	// mongoURI := fmt.Sprintf("mongodb+srv://%s:<%s>@cluster0-zz20n.mongodb.net/test?retryWrites=true&w=majority", "icecreamhotz", "zazaza1b")
	// app, err := InitialApplication(mongoURI, 1*time.Second)
	// if err != nil {
	// 	log.Fatal("failed to create event: %s\n", err)
	// }

	databaseErr := app.DB.Ping(context.TODO(), nil)

	if databaseErr != nil {
		log.Fatal(databaseErr)
	}

	router := gin.Default()

	userRoute := router.Group("/user")
	routes.UserRoute(userRoute)

	router.Run(":" + os.Getenv("port"))
}
