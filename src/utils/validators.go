package utils

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

func AsFieldError(err error, target *validator.ValidationErrors) bool {
	return errors.As(err, target)
}

func UnmarshalTypeErrorAs(err error, target **json.UnmarshalTypeError) bool {
	return errors.As(err, target)
}
