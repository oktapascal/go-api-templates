package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-rental/domain"
	"net/http"
)

type Handler struct {
	svc      domain.UserService
	validate *validator.Validate
}

func (hdl *Handler) Store() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user = new(domain.User)

		var decoder = json.NewDecoder(request.Body)
		var err = decoder.Decode(&user)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(user)
		if err != nil {
			panic(err)
		}

		hdl.svc.Save(request.Context(), user)
	}
}

func (hdl *Handler) GetByEmail() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
