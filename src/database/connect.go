package database

import (
	"fmt"

	"github.com/icecreamhotz/movie-ticket/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func NewDatabase(config *configs.ConfigDatabase) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.URL,
	))
	if err != nil {
		fmt.Println("connect database error : ", err)
		return nil, err
	}
	return client, nil
}
