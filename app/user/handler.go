package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-rental/domain"
	"go-rental/exceptions"
	"go-rental/response"
	"net/http"
)

type Handler struct {
	svc      domain.UserService
	validate *validator.Validate
}

func (hdl *Handler) Store() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user = domain.User{}

		var decoder = json.NewDecoder(request.Body)
		var err = decoder.Decode(&user)
		if err != nil {
			exceptions.InternalServerHandler(writer, err)
		}

		err = hdl.validate.Struct(user)

		if err != nil {
			var formatErrors = exceptions.BadRequestHandler(err)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)

			var responseError = response.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: formatErrors,
			}

			encoder := json.NewEncoder(writer)

			err = encoder.Encode(responseError)
			if err != nil {
				exceptions.InternalServerHandler(writer, err)
			}

			return
		}
	}
}

func (hdl *Handler) GetByEmail() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
