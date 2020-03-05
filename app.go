package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/icecreamhotz/movie-ticket/controllers"
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
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongodb:27017/", "root", "root")
	kk, err := InitialApplication(mongoURI, 1*time.Second)
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e := kk.Handler.DB.Ping(context.TODO(), nil)

	if e != nil {
		log.Fatal(e)
	}
	log.Fatal("Connect success")
	// kk.Handler.UsersPost()
	// mongoURI := fmt.Sprintf("mongodb+srv://%s:<%s>@cluster0-zz20n.mongodb.net/test?retryWrites=true&w=majority", "icecreamhotz", "zazaza1b")
	// conn, err := database.ConnectDatabase(mongoURI, 10*time.Second)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
