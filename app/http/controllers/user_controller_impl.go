package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	exceptions2 "go-rental/app/exceptions"
	"go-rental/app/http/requests"
	"go-rental/app/services"
	"go-rental/responses"
	"net/http"
)

type UserControllerImpl struct {
	Validate    *validator.Validate
	UserService services.UserService
}

func NewUserController(validate *validator.Validate, userService services.UserService) UserController {
	return &UserControllerImpl{Validate: validate, UserService: userService}
}

func (controller *UserControllerImpl) Store(writer http.ResponseWriter, request *http.Request) {
	userCreateRequest := requests.UserCreateRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		exceptions2.InternalServerHandler(writer, err)
	}

	err = controller.Validate.Struct(userCreateRequest)
	if err != nil {
		formatErrors := exceptions2.BadRequestHandler(err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		responseError := responses.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Errors: formatErrors,
		}

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(responseError)
		if err != nil {
			exceptions2.InternalServerHandler(writer, err)
		}

		return
	}
	//controller.UserService.Save(request.Context(), userCreateRequest)
}

func (controller *UserControllerImpl) GetByEmail(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
