package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationReporer interface {
	GetID(string, string) (Reservation, error)
	Create(*Reservation, *Movie) (primitive.ObjectID, error)
	GetAll(string) ([]*Reservation, error)
}

type ReservationRepository struct {
	DB *mongo.Collection
}

type Reservation struct {
	MovieId     string    `form:"movieId" json:"movieId" binding:"required"`
	SeatNo      []string  `form:"seatNo" json:"seatNo" binding:"required"`
	UserId      string    `form:"userId" json:"userId" binding:"required"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	MovieDetail *Movie    `json:"movieDetail,omitempty"`
}

func (repo *ReservationRepository) GetID(id, userId string) (Reservation, error) {
	var reservation Reservation
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "userId": userId}
	err := repo.DB.FindOne(context.TODO(), filter).Decode(&reservation)
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

func (repo *ReservationRepository) GetAll(userId string) ([]*Reservation, error) {
	cur, err := repo.DB.Find(context.TODO(), bson.M{
		"userId": userId,
	})
	if err != nil {
		return nil, err
	}
	var results []*Reservation
	for cur.Next(context.TODO()) {
		var elem *Reservation
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}

		results = append(results, elem)

	}
	return results, nil
}

func (repo *ReservationRepository) Create(reservation *Reservation, movie *Movie) (primitive.ObjectID, error) {
	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"movieId":     reservation.MovieId,
		"seatNo":      reservation.SeatNo,
		"userId":      reservation.UserId,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"movieDetail": movie,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
