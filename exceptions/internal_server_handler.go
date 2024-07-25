package exceptions

import (
	"encoding/json"
	"go-rental/libs"
	"go-rental/responses"
	"net/http"
)

// InternalServerHandler is a function that handles HTTP 500 Internal Server Error responses.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error and panics.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - error: The error interface containing the details of the error.
//
// Return:
// This function does not return any value.
func InternalServerHandler(writer http.ResponseWriter, error error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := responses.ErrorResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: error,
	}

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		libs.CreateLoggerFile().Panic(err)
	}

	libs.CreateLoggerFile().Panic(error)
}
