package controllers

import "net/http"

type UserController interface {
	Store(writer http.ResponseWriter, request *http.Request)
	GetByEmail(writer http.ResponseWriter, request *http.Request)
}
