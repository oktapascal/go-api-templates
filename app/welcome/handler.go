package welcome

import (
	"github.com/spf13/viper"
	"go-rental/exceptions"
	"net/http"
)

type Handler struct {
}

// Welcome is a handler function that writes a welcome message to the HTTP response.
//
// It takes a http.ResponseWriter and a *http.Request as parameters.
// It does not return anything.
func (hdl *Handler) Welcome() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
		if err != nil {
			exceptions.InternalServerHandler(writer, err)
		}
	}
}
