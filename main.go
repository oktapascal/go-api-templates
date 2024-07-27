package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go-rental/app/user"
	"go-rental/app/welcome"
	"go-rental/config"
	"go-rental/middlewares"
	"net/http"
)

// main is the entry point for the application.
// It does not take any parameters and does not return any value.
func main() {
	config.InitConfig()

	var db, err = config.ConnectDatabase()
	if err != nil {
		config.CreateLoggerFile().Fatal(err)
	}

	var validate = config.CreateValidator()

	var router = chi.NewRouter()

	router.Use(middlewares.LoggerMiddleware)
	router.Use(middlewares.RecoverMiddleware)

	var welcomeHandler = welcome.Wire()
	var userHandler = user.Wire(validate, db)

	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", welcomeHandler.Welcome())
	router.Post("/user", userHandler.Store())

	config.CreateLoggerConsole(nil).Info("Application Started")
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		config.CreateLoggerFile().Fatal(err)
	}
}
