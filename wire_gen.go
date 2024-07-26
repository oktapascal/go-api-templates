// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"go-rental/app/http/controllers"
	"go-rental/app/services"
	"go-rental/config"
)

// Injectors from init_wire.go:

func InitializeWelcomeController() controllers.WelcomeController {
	welcomeController := controllers.NewWelcomeController()
	return welcomeController
}

func InitializeUserController() controllers.UserController {
	validate := config.CreateValidator()
	userService := services.NewUserService()
	userController := controllers.NewUserController(validate, userService)
	return userController
}

// init_wire.go:

var ProviderSet = wire.NewSet(config.CreateValidator)
