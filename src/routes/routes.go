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
	helloController := controllers.NewHelloController()

	r := gin.Default()
	r.POST("/person", personController.Create)
	r.GET("/", helloController.Hello)

	r.Run()
}
