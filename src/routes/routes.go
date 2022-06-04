package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuellfa/registerserver/src/controllers"
	"github.com/samuellfa/registerserver/src/database"
	"github.com/samuellfa/registerserver/src/services"
)

func HandleRequest() {
	personService := services.NewPersonService(database.DB)
	personController := controllers.NewPersonController(personService)

	r := gin.Default()
	r.POST("/person", personController.Create)

	r.Run()
}
