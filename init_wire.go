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

func InitializeUserService() services.UserService {
	wire.Build(services.NewUserService)
	return nil
}

var UserSet = wire.NewSet(
	controllers.NewUserController,
	ProviderSet,
	services.NewUserService,
)

func InitializeUserController() controllers.UserController {
	wire.Build(UserSet)
	return nil
}
