package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikosmpi/gorestapi/db"
	"github.com/nikosmpi/gorestapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8181")
}
