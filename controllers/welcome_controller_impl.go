package controllers

import (
	"github.com/spf13/viper"
	"go-rental/exceptions"
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
		exceptions.InternalServerHandler(writer, err)
	}
}
