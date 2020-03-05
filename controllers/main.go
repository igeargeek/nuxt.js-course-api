package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceHandler struct {
	DB *mongo.Client
}

func NewHandler(db *mongo.Client) *ServiceHandler {
	return &ServiceHandler{
		DB: db,
	}
}
