package errors

import (
	"net/http"
)

type ValidationErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	messages []ValidationErrorField
}

func (e *ValidationError) Error() interface{} {
	return e.messages
}

func (e *ValidationError) HttpStatus() int {
	return http.StatusBadRequest
}

func (e *ValidationError) AddError(err ValidationErrorField) {
	e.messages = append(e.messages, err)
}

func (e *ValidationError) IsEmpty() bool {
	return len(e.messages) == 0
}
