package models

import (
	"context"
	"mime/multipart"
	"time"

	"app/src/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserReporer interface {
	GetID(ID primitive.ObjectID) (User, error)
	Create(user *User) (primitive.ObjectID, error)
	CreateRefreshToken(token *RefreshToken) error
	FindByUsername(username string) (User, error)
	CheckUserNameExists(username string) (primitive.ObjectID, error)
	GetAccessToken(token string) (AccessToken, error)
	GetRefreshToken(token string) (RefreshToken, error)
	RemoveRefreshToken(ID primitive.ObjectID) error
}

type UserRepository struct {
	DB             *mongo.Collection
	DBAccessToken  *mongo.Collection
	DBRefreshToken *mongo.Collection
}

type AccessToken struct {
	ID     primitive.ObjectID `bson:"_id",omitempty`
	UserID primitive.ObjectID `bson:"userId",omitempty`
	Token  string             `bson:"accessToken"`
}

type RefreshToken struct {
	ID     primitive.ObjectID `bson:"_id",omitempty`
	UserID primitive.ObjectID `bson:"userId",omitempty`
	Token  string             `bson:"refreshToken"`
}

type User struct {
	ID         primitive.ObjectID    `bson:"_id",omitempty`
	Name       string                `bson:"name" form:"name" json:"name" binding:"required"`
	Username   string                `bson:"username" form:"username" json:"username" binding:"required"`
	Password   string                `bson:"password" form:"password" json:"password" binding:"required"`
	AvatarFile *multipart.FileHeader `form:"avatar"`
	Avatar     string                `bson:"avatar" json:"avatar"`
	CreatedAt  time.Time             `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time             `bson:"updatedAt" json:"updatedAt"`
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

func (repo *UserRepository) CheckUserNameExists(username string) (primitive.ObjectID, error) {
	_, err := repo.FindByUsername(username)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return primitive.NilObjectID, err
		}
		return primitive.NilObjectID, nil
	}
	return primitive.NilObjectID, constants.ErrRowExists
}

func (repo *UserRepository) Create(user *User) (primitive.ObjectID, error) {
	timeNow := time.Now()

	res, err := repo.DB.InsertOne(context.TODO(), bson.M{
		"name":      user.Name,
		"username":  user.Username,
		"password":  user.Password,
		"avatar":    "public/" + user.Avatar,
		"createdAt": timeNow,
		"updatedAt": timeNow,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo *UserRepository) CreateRefreshToken(token *RefreshToken) error {
	_, err := repo.DBRefreshToken.InsertOne(context.TODO(), bson.M{
		"userId":       token.UserID,
		"refreshToken": token.Token,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetAccessToken(token string) (AccessToken, error) {
	var accessToken AccessToken
	err := repo.DBAccessToken.FindOne(context.TODO(), bson.D{{"accessToken", token}}).Decode(&accessToken)
	if err == nil {
		return accessToken, nil
	}
	return accessToken, err
}

func (repo *UserRepository) GetRefreshToken(token string) (RefreshToken, error) {
	var refreshToken RefreshToken
	err := repo.DBRefreshToken.FindOne(context.TODO(), bson.D{{"refreshToken", token}}).Decode(&refreshToken)
	if err == nil {
		return refreshToken, nil
	}
	return refreshToken, err
}

func (repo *UserRepository) RemoveRefreshToken(id primitive.ObjectID) error {
	_, err := repo.DBRefreshToken.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	return err
}
