package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuellfa/registerserver/src/controllers"
	"github.com/samuellfa/registerserver/src/database"
)

func HandleRequest() {
	personController := controllers.NewPersonController(database.DB)

	r := gin.Default()
	r.POST("/person", personController.Create)

	r.Run()
}
