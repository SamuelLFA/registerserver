package errors

import (
	"fmt"
	"net/http"
)

type ValidationError struct {
	messages []string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%q", e.messages)
}

func (e *ValidationError) HttpStatus() int {
	return http.StatusBadRequest
}

func (e *ValidationError) AddError(err string) {
	e.messages = append(e.messages, err)
}

func (e *ValidationError) IsEmpty() bool {
	return len(e.messages) == 0
}
