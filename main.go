package main

import (
	"chaitanyaallu.dev/event-management/db"
	"chaitanyaallu.dev/event-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	serverEngine := gin.Default()
	routes.RegisterRoutes(serverEngine)
	serverEngine.Run(":8080")
}
