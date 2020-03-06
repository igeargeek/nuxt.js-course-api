// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"time"

	"github.com/icecreamhotz/movie-ticket/configs"
	"github.com/icecreamhotz/movie-ticket/controllers"
	"github.com/icecreamhotz/movie-ticket/database"
	"github.com/icecreamhotz/movie-ticket/models"
	"github.com/icecreamhotz/movie-ticket/utils"
)

// Injectors from wire.go:

func InitialApplication(mongoURI string, timeout time.Duration) (App, error) {
	configDatabase := configs.NewConfig(mongoURI, timeout)
	client, err := database.NewDatabase(configDatabase)
	if err != nil {
		return App{}, err
	}
	userRepository := models.NewUserRepository(client)
	translator := utils.NewValidateTranslation()
	userHandler := controllers.NewUserHandler(userRepository, translator)
	app := NewAppDatabase(client, userHandler)
	return app, nil
}
