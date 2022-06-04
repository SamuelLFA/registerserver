package helpers

import (
	"github.com/go-playground/validator/v10"
	e "github.com/samuellfa/registerserver/src/errors"
)

func CreateValidationErrors(errors validator.ValidationErrors) e.ValidationError {
	var errorValidation e.ValidationError
	for _, err := range errors {
		errorValidation.AddError(e.ValidationErrorField{Field: err.Field(), Message: err.Tag()})
	}

	return errorValidation
}
