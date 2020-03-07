package models

import (
	"context"
	"time"

	"app/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserReporer interface {
	GetID(ID primitive.ObjectID) (User, error)
	Create(user *User) (primitive.ObjectID, error)
	FindByUsername(username string) (User, error)
}

type UserRepository struct {
	DB *mongo.Collection
}

type User struct {
	ID        primitive.ObjectID `bson:"_id",omitempty`
	Name      string             `bson:"name" form:"name" json:"name" binding:"required"`
	Username  string             `bson:"username" form:"username" json:"username" binding:"required"`
	Password  string             `bson:"password" form:"password" json:"password" binding:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (repo *UserRepository) GetID(ID primitive.ObjectID) (User, error) {
	var user User
	err := repo.DB.FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&user)
	if err == nil {
		return user, nil
	}
	return user, nil
}

func (repo *UserRepository) FindByUsername(username string) (User, error) {
	var user User
	err := repo.DB.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err == nil {
		return user, nil
	}
	return user, err
}

func (repo *UserRepository) Create(user *User) (primitive.ObjectID, error) {
	_, err := repo.FindByUsername(user.Username)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return primitive.NilObjectID, err
		}
	} else {
		return primitive.NilObjectID, utils.ErrRowExists
	}

	timeNow := time.Now()

	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name":      user.Name,
		"username":  user.Username,
		"password":  user.Password,
		"createdAt": timeNow,
		"updatedAt": timeNow,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
