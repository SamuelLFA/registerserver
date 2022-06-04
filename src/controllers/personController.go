package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/samuellfa/registerserver/src/models"
	"gorm.io/gorm"
)

type personController struct {
	DB *gorm.DB
}

func NewPersonController(db *gorm.DB) personController {
	return personController{db}
}

func (controller *personController) Create(ctx *gin.Context) {
	var person models.Person

	if err := ctx.ShouldBindWith(&person, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.DB.Save(&person)
	ctx.JSON(http.StatusCreated, person)
}
