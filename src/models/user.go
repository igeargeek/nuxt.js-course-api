package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserReporer interface {
	GetID(ID int) (int, error)
	Create(user *User) (primitive.ObjectID, error)
}

type UserRepository struct {
	DB *mongo.Client
}

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (repo *UserRepository) GetID(ID int) (int, error) {
	return 0, nil
}

func (repo *UserRepository) Create(user *User) (primitive.ObjectID, error) {
	collection := repo.DB.Database("movie_ticket").Collection("users")
	res, err := collection.InsertOne(context.TODO(), bson.M{
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
