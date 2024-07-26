package controllers

import "net/http"

type WelcomeController interface {
	Welcome(writer http.ResponseWriter, request *http.Request)
}
