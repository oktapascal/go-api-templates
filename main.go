package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-rental/controllers"
	"go-rental/libs"
	"go-rental/middlewares"
	"go-rental/services"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/probandula/figlet4go"
	"github.com/spf13/viper"
)

// main is the entry point for the application.
// It does not take any parameters and does not return any value.
func main() {
	ascii := figlet4go.NewAsciiRender()
	optionAscii := figlet4go.NewRenderOptions()
	optionAscii.FontName = "standard"

	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	router := chi.NewRouter()

	router.Use(middlewares.LoggerMiddleware)
	router.Use(middleware.Recoverer)

	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		libs.CreateLoggerFile().Fatal(err)
		return
	}

	welcomeController := controllers.NewWelcomeController()
	userService := services.NewUserService()
	userController := controllers.NewUserController(validate, userService)
	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", welcomeController.Welcome)
	router.Post("/user", userController.Store)

	banner, _ := ascii.RenderOpts("RW"+"v"+viper.GetString("APP_VERSION"), optionAscii)
	fmt.Print(banner)

	libs.CreateLoggerConsole(nil).Info("Application Started")
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		libs.CreateLoggerFile().Fatal(err)
	}
}
