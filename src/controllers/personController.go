package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/samuellfa/registerserver/src/controllers/dto"
	"github.com/samuellfa/registerserver/src/services"
)

type personController struct {
	service *services.PersonService
}

func NewPersonController(service *services.PersonService) personController {
	return personController{service: service}
}

func (controller *personController) Create(ctx *gin.Context) {
	var request dto.CreatePersonRequest

	if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := controller.service.Create(request, ctx)
	if err != nil {
		ctx.JSON(err.HttpStatus(), gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.NewCreatePersonResponse(*person)
	ctx.JSON(http.StatusCreated, response)
}
