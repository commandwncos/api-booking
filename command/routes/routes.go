package routes

import (
	"github.com/commandwncos/api-booking/command/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/signup", signup)
	server.POST("/login", login)

	auth := server.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	auth.GET("/events", HandleGetEvents)
	auth.GET("/events/:id", HandleGetEventById)
	auth.POST("/events", HandlePostEvents)
	auth.PUT("/events/:id", HandleUpdateEventById)
	auth.DELETE("/events/:id", HandleDeleteEventById)

	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register", cancelForEvent)

}
