package database

import (
	"log"
	"os"

	"app/src/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func NewDatabase(config configs.ConfigDatabase) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.URL,
	))

	if err != nil {
		log.Fatal("connect database error : ", err)
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("ping error", err)
	}

	return client.Database(os.Getenv("MONGO_DBNAME"))
}
