package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieReporer interface {
	GetID(string) (Movie, error)
	Create(*Movie) (primitive.ObjectID, error)
	DeleteID(string) error
	Edit(string, *Movie) error
	GetAll() ([]*Movie, error)
}

type MovieRepository struct {
	DB *mongo.Collection
}

type Movie struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `form:"name" json:"name" binding:"required"`
	Genre        []string           `form:"genre" json:"genre" binding:"required"`
	PosterURL    string             `form:"posterUrl" json:"posterUrl" binding:"required"`
	YoutubeURL   string             `form:"youtubeUrl" json:"youtubeUrl" binding:"required"`
	Description  string             `form:"description" json:"description" binding:"required"`
	Duration     int                `form:"duration" json:"duration" binding:"required"`
	ReservedSeat []string           `form:"reservedSeat" json:"reservedSeat"`
	ShowingAt    string             `form:"showingAt" json:"showingAt" binding:"required"`
}

func (repo *MovieRepository) GetID(id string) (Movie, error) {
	var movie Movie
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := repo.DB.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (repo *MovieRepository) GetAll() ([]*Movie, error) {
	cur, err := repo.DB.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*Movie
	for cur.Next(context.TODO()) {
		var elem *Movie
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}

		results = append(results, elem)

	}
	return results, nil
}

func (repo *MovieRepository) Create(movie *Movie) (primitive.ObjectID, error) {
	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name":        movie.Name,
		"genre":       movie.Genre,
		"posterUrl":   movie.PosterURL,
		"youtubeUrl":  movie.YoutubeURL,
		"description": movie.Description,
		"duration":    movie.Duration,
		"showingAt":   movie.ShowingAt,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo *MovieRepository) Edit(id string, movie *Movie) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := repo.DB.FindOneAndUpdate(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"name":         movie.Name,
			"genre":        movie.Genre,
			"posterUrl":    movie.PosterURL,
			"youtubeUrl":   movie.YoutubeURL,
			"description":  movie.Description,
			"duration":     movie.Duration,
			"reservedSeat": movie.ReservedSeat,
			"showingAt":    movie.ShowingAt,
		},
	})
	if err != nil {
		return err.Err()
	}

	return nil
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
