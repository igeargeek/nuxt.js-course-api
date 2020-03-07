package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieReporer interface {
	GetID(ID int) (int, error)
	Create(movie *Movie) (primitive.ObjectID, error)
}

type MovieRepository struct {
	DB *mongo.Collection
}

type Movie struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func (repo *MovieRepository) Create(movie *Movie) (primitive.ObjectID, error) {
	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name": movie.Name,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo *MovieRepository) GetID(number int) (int, error) {
	return 0, nil
}
