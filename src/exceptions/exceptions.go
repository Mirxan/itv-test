package exceptions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mirxan/itv-test/utils"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewBadRequestError(message string) *AppError {
	return &AppError{Code: 400, Message: message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: 404, Message: message}
}

func NewValidationErrors(err error) map[string]string {
	errorMessages := make(map[string]string)

	var verrs validator.ValidationErrors
	if ok := utils.AsFieldError(err, &verrs); ok {
		for _, err := range verrs {
			switch err.Tag() {
			case "oneof":
				errorMessages[err.Field()] = fmt.Sprintf("%s should be %s", err.Field(), strings.Join(strings.Split(err.Param(), " "), " or "))
			case "form_required", "required":
				errorMessages[err.Field()] = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errorMessages[err.Field()] = fmt.Sprintf("%s must be a valid email address", err.Field())
			case "min":
				errorMessages[err.Field()] = fmt.Sprintf("%s must have at least %s characters", err.Field(), err.Param())
			case "max":
				errorMessages[err.Field()] = fmt.Sprintf("%s must have at most %s characters", err.Field(), err.Param())
			case "gt":
				errorMessages[err.Field()] = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
			case "gte":
				errorMessages[err.Field()] = fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
			default:
				errorMessages[err.Field()] = fmt.Sprintf("Validation failed on field %s", err.Field())
			}
		}
		return errorMessages
	}

	var unmarshalErr *json.UnmarshalTypeError
	if ok := utils.UnmarshalTypeErrorAs(err, &unmarshalErr); ok {
		errorMessages[unmarshalErr.Field] = fmt.Sprintf("Invalid type for field '%s': expected %s", unmarshalErr.Field, unmarshalErr.Type.String())
		return errorMessages
	}

	errorMessages["error"] = err.Error()
	return errorMessages
}
