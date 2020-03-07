package models

import (
	"context"

	"app/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserReporer interface {
	GetID(ID int) (int, error)
	Create(user *User) (primitive.ObjectID, error)
}

type UserRepository struct {
	DB *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id",omitempty`
	Name     string             `bson:"userId" form:"name" json:"name" binding:"required"`
	Username string             `bson:"url" form:"username" json:"username" binding:"required"`
	Password string             `bson:"title" form:"password" json:"password" binding:"required"`
}

func (repo *UserRepository) GetID(ID int) (int, error) {
	return 0, nil
}

func (repo *UserRepository) Create(user *User) (primitive.ObjectID, error) {
	var result bson.M
	err := repo.DB.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if err == nil {
		return primitive.NilObjectID, utils.ErrRowExists
	}

	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
