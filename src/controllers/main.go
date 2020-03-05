package controllers

import (
	ut "github.com/go-playground/universal-translator"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceHandler struct {
	DB        *mongo.Client
	Validator ut.Translator
}

func NewHandler(db *mongo.Client, validator ut.Translator) *ServiceHandler {
	return &ServiceHandler{
		DB:        db,
		Validator: validator,
	}
}
