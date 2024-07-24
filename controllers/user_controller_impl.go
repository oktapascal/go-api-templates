package controllers

import (
	"encoding/json"
	"go-rental/requests"
	"go-rental/services"
	"log"
	"net/http"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Store(writer http.ResponseWriter, request *http.Request) {
	userCreateRequest := requests.UserCreateRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		log.Panic(err)
	}

	controller.UserService.Save(request.Context(), userCreateRequest)
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) GetByEmail(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
