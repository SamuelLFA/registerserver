package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/samuellfa/registerserver/src/controllers"
	d "github.com/samuellfa/registerserver/src/database"
	s "github.com/samuellfa/registerserver/src/services"
)

func HandleRequest() {
	personService := s.NewPersonService(d.DB)
	personController := c.NewPersonController(personService)
	helloController := c.NewHelloController()

	r := gin.Default()
	r.POST("/person", personController.Create)
	r.GET("/", helloController.Hello)

	r.Run()
}
