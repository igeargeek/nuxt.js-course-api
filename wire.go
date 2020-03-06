// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"time"

	"app/src/configs"
	"app/src/controllers"
	"app/src/database"
	"app/src/models"
	"app/src/utils"

	"github.com/google/wire"
)

func InitialApplication(mongoURI string, timeout time.Duration) (App, error) {
	wire.Build(configs.NewConfig, database.NewDatabase, models.NewUserRepository, utils.NewValidateTranslation, controllers.NewUserHandler, NewAppDatabase)
	return App{}, nil
}
