package controllers

import (
	"github.com/spf13/viper"
	"go-rental/libs"
	"net/http"
)

type WelcomeControllerImpl struct {
}

func NewWelcomeController() WelcomeController {
	return &WelcomeControllerImpl{}
}

func (controller WelcomeControllerImpl) Welcome(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
	if err != nil {
		http.Error(writer, http.StatusText(500), 500)
		libs.CreateLogEntry(request).Error(err)
		panic(err)
	}
}
