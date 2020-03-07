package models

import "go.mongodb.org/mongo-driver/mongo"

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{
		DB: db.Collection("users"),
	}
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return MovieRepository{
		DB: db.Collection("movies"),
	}
}
