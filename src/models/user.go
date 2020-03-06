package models

import "go.mongodb.org/mongo-driver/mongo"

type UserReporer interface {
	GetID(ID int) (int, error)
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
