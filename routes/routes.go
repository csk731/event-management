package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(serverEngine *gin.Engine) {
	serverEngine.GET("/events", getEvents)
	serverEngine.GET("/events/:id", getEvent)
	serverEngine.POST("/events", createEvent)
	serverEngine.PUT("/events/:id", updateEvent)
	serverEngine.DELETE("/events/:id", deleteEvent)

	serverEngine.POST("/signup", signUp)
	serverEngine.POST("/login", login)
	serverEngine.GET("/users/:id", getUserByID)

}