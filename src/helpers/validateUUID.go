package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	e "github.com/samuellfa/registerserver/src/errors"
)

func GetIdFromContext(ctx *gin.Context) (string, e.APIError) {
	id, idIsPresent := ctx.Params.Get("id")
	if !idIsPresent {
		var errors e.ValidationError
		err := e.ValidationErrorField{
			Field:   "id",
			Message: "required",
		}
		errors.AddError(err)
		return "", &errors
	}
	_, err := uuid.Parse(id)
	if err != nil {
		var errors e.ValidationError
		err := e.ValidationErrorField{
			Field:   "id",
			Message: "invalid uuid",
		}
		errors.AddError(err)
		return "", &errors
	}
	return id, nil
}
