package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieReporer interface {
	GetID(string) (Movie, error)
	Create(movie *Movie) (primitive.ObjectID, error)
	DeleteID(string) error
}

type MovieRepository struct {
	DB *mongo.Collection
}

type Movie struct {
	Name         string   `form:"name" json:"name" binding:"required"`
	Genre        []string `form:"genre" json:"genre" binding:"required"`
	PosterURL    string   `form:"posterUrl" json:"posterUrl" binding:"required"`
	YoutubeURL   string   `form:"youtubeUrl" json:"youtubeUrl" binding:"required"`
	Description  string   `form:"description" json:"description" binding:"required"`
	Duration     int      `form:"duration" json:"duration" binding:"required"`
	ReservedSeat []string `form:"reservedSeat" json:"reservedSeat"`
}

func (repo *MovieRepository) GetID(id string) (Movie, error) {
	var movie Movie
	collection := repo.DB.Database("movie_ticket").Collection("movies")
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (repo *MovieRepository) Create(movie *Movie) (primitive.ObjectID, error) {
	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name":        movie.Name,
		"genre":       movie.Genre,
		"posterUrl":   movie.PosterURL,
		"youtubeUrl":  movie.YoutubeURL,
		"description": movie.Description,
		"duration":    movie.Duration,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo *MovieRepository) DeleteID(id string) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	_, err := repo.DB.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
