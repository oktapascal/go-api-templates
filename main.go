package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"go-rental/app/http/middlewares"
	"go-rental/config"
	"net/http"
)

// main is the entry point for the application.
// It does not take any parameters and does not return any value.
func main() {
	config.InitConfig()

	router := chi.NewRouter()

	router.Use(middlewares.LoggerMiddleware)
	router.Use(middleware.Recoverer)

	welcomeController := InitializeWelcomeController()
	userController := InitializeUserController()
	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", welcomeController.Welcome)
	router.Post("/user", userController.Store)

	config.CreateLoggerConsole(nil).Info("Application Started")
	err := http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		config.CreateLoggerFile().Fatal(err)
	}
}
