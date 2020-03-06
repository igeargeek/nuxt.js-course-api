package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"app/src/controllers"
	"app/src/routes"

	"github.com/gin-gonic/gin"
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
	mongoURI := fmt.Sprintf(os.Getenv("mongo_uri"))
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
