// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"time"

	"github.com/google/wire"
	"github.com/icecreamhotz/movie-ticket/configs"
	"github.com/icecreamhotz/movie-ticket/controllers"
	"github.com/icecreamhotz/movie-ticket/database"
	"github.com/icecreamhotz/movie-ticket/models"
	"github.com/icecreamhotz/movie-ticket/utils"
)

func InitialApplication(mongoURI string, timeout time.Duration) (App, error) {
	wire.Build(configs.NewConfig, database.NewDatabase, models.NewUserRepository, utils.NewValidateTranslation, controllers.NewUserHandler, NewAppDatabase)
	return App{}, nil
}
