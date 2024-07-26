package welcome

import (
	"github.com/spf13/viper"
	"go-rental/exceptions"
	"net/http"
)

type Handler struct {
}

func (hdl *Handler) Welcome() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
		if err != nil {
			exceptions.InternalServerHandler(writer, err)
		}
	}
}
