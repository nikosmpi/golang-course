package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikosmpi/gorestapi/db"
	"github.com/nikosmpi/gorestapi/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8181")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   event,
	})
	//events := models.GetAllEvents()
	//context.JSON(http.StatusOK, events)
}
