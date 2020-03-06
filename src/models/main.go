package models

import "go.mongodb.org/mongo-driver/mongo"

func NewUserRepository(db *mongo.Client) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func NewMovieRepository(db *mongo.Client) MovieRepository {
	return MovieRepository{
		DB: db,
	}
}
