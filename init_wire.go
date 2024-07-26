//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-rental/app/http/controllers"
	"go-rental/app/services"
	"go-rental/config"
)

var ProviderSet = wire.NewSet(config.CreateValidator)

func InitializeWelcomeController() controllers.WelcomeController {
	wire.Build(controllers.NewWelcomeController)
	return nil
}

func InitializeUserController() controllers.UserController {
	wire.Build(controllers.NewUserController, ProviderSet, services.NewUserService)
	return nil
}
