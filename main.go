package main

import (
	"fmt"
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

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
	})

	banner, _ := ascii.RenderOpts("RW"+"v"+viper.GetString("APP_VERSION"), optionAscii)
	fmt.Print(banner)

	http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
}
