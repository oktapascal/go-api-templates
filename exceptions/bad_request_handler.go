package exceptions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-rental/response"
	"net/http"
)

type formatError struct {
	Param   string
	Message string
}

// convertTagToMessage converts a validator.FieldError's tag into a human-readable error message.
// It supports the following tags: "required", "email", "min", and "max".
// For "min" and "max" tags, it includes the field's parameter in the error message.
// If the tag is not recognized, it returns the original error message.
func convertTagToMessage(ex validator.FieldError) string {
	switch ex.Tag() {
	case "required":
		return "kolom ini tidak boleh kosong"
	case "email":
		return "email tidak valid"
	case "min":
		return fmt.Sprint("kolom ini harus memiliki panjang minimal ", ex.Param(), " karakter")
	case "max":
		return fmt.Sprint("kolom ini harus memiliki panjang maksimal ", ex.Param(), " karakter")
	default:
		return ex.Error()
	}
}

// FormatErrors is a function that processes a validation error and returns a slice of formatError.
// It is designed to handle errors generated by the go-playground/validator/v10 package.
// If the input error is a validator.ValidationErrors, it extracts field-level validation errors,
// converts them into formatError structs using the convertTagToMessage function, and returns the slice.
// If the input error is not a validator.ValidationErrors, it returns nil.
func FormatErrors(error error) []formatError {
	var exception validator.ValidationErrors

	// Check if the input error is a validator.ValidationErrors
	if errors.As(error, &exception) {
		// Create a slice to store the formatError structs
		fieldErrors := make([]formatError, len(exception))

		// Iterate over the validation errors and convert them into formatError structs
		for index, ex := range exception {
			fieldErrors[index] = formatError{
				Param:   ex.Field(),
				Message: convertTagToMessage(ex),
			}
		}

		// Return the slice of formatError structs
		return fieldErrors
	}

	// If the input error is not a validator.ValidationErrors, return nil
	return nil
}

// BadRequestHandler handles HTTP 400 Bad Request responses.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it calls the InternalServerHandler function.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - error: The error interface containing the details of the error.
func BadRequestHandler(writer http.ResponseWriter, error any) {
	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Bad Request
	writer.WriteHeader(http.StatusBadRequest)

	// Create an error response with the status code and error details
	responseError := response.ErrorResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: error,
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if err := encoder.Encode(responseError); err != nil {
		// If there was an error, call the InternalServerHandler function
		InternalServerHandler(writer, err)
	}
}
