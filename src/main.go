package main

import (
	"github.com/samuellfa/registerserver/src/database"
	"github.com/samuellfa/registerserver/src/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.HandleRequest()
}
