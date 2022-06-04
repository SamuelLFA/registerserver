package errors

import (
	"net/http"
)

type internalServerError struct {
	message string
}

func NewInternalServerError(message string) *internalServerError {
	return &internalServerError{message}
}

func (err *internalServerError) Error() interface{} {
	return err.message
}

func (err *internalServerError) HttpStatus() int {
	return http.StatusInternalServerError
}
