package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-rental/controllers"
	"go-rental/services"
	"log"
	"net/http"

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
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	userService := services.NewUserService(validate)
	userController := controllers.NewUserController(userService)
	//router.Use(middlewares.AuthorizationCheckMiddleware)
	//router.Use(middlewares.VerifyTokenMiddleware)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Panic(err)
		}
	})

	router.Post("/user", userController.Store)

	banner, _ := ascii.RenderOpts("RW"+"v"+viper.GetString("APP_VERSION"), optionAscii)
	fmt.Print(banner)

	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
