package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type helloController struct {
}

func NewHelloController() helloController {
	return helloController{}
}

func (controller *helloController) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"mensagem": "Pq és a mulher mais linda desse mundo?",
	})
}
