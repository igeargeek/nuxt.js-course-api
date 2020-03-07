package models

import "go.mongodb.org/mongo-driver/mongo"

func NewUserRepository(db *mongo.Database) UserReporer {
	return &UserRepository{
		DB: db.Collection("users"),
	}
}

func NewMovieRepository(db *mongo.Database) MovieReporer {
	return &MovieRepository{
		DB: db.Collection("movies"),
	}
}
