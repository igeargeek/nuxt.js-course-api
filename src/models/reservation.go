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
	Create(*Reservation) (primitive.ObjectID, error)
	// DeleteID(string) error
	// Edit(string, *Movie) error
	// GetAll() ([]*Movie, error)
}

type ReservationRepository struct {
	DB *mongo.Collection
}

type Reservation struct {
	MovieId   string    `form:"movieId" json:"movieId" binding:"required"`
	SeatNo    []string  `form:"seatNo" json:"seatNo" binding:"required"`
	UserId    string    `form:"userId" json:"userId" binding:"required"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
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

// func (repo *MovieRepository) GetAll() ([]*Movie, error) {
// 	cur, err := repo.DB.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	var results []*Movie
// 	for cur.Next(context.TODO()) {
// 		var elem *Movie
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			return results, err
// 		}

// 		results = append(results, elem)

// 	}
// 	return results, nil
// }

func (repo *ReservationRepository) Create(reservation *Reservation) (primitive.ObjectID, error) {
	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"movieId":    reservation.MovieId,
		"seatNo":     reservation.SeatNo,
		"userId":     reservation.UserId,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
