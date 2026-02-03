package main

import (
	"github.com/commandwncos/api-booking/command/private/database"
	"github.com/commandwncos/api-booking/command/routes"
	g "github.com/gin-gonic/gin"
)

func main() {
	connex := database.Connect()
	defer connex.Close()
	server := g.Default()
	routes.RegisterRoutes(server)
	server.Run(":8000") //localhost or 127.0.0.1
}
