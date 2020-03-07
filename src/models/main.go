package models

import "go.mongodb.org/mongo-driver/mongo"

func NewUserRepository(db *mongo.Database) UserReporer {
	return &UserRepository{
		DB:             db.Collection("users"),
		DBAccessToken:  db.Collection("access_token"),
		DBRefreshToken: db.Collection("refresh_token"),
	}
}

func NewMovieRepository(db *mongo.Database) MovieReporer {
	return &MovieRepository{
		DB: db.Collection("movies"),
	}
}

func NewReservationRepository(db *mongo.Database) ReservationReporer {
	return &ReservationRepository{
		DB: db.Collection("reservations"),
	}
}
