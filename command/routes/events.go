package routes

import (
	"net/http"

	"github.com/commandwncos/api-booking/command/private/database"
	"github.com/commandwncos/api-booking/command/utils"
	"github.com/commandwncos/api-booking/models"
	"github.com/gin-gonic/gin"
)

// GET /events
func HandleGetEvents(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	events, err := models.GetEventsByUser(database.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

// GET /events/:id
func HandleGetEventById(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventID, err := utils.ParseToInteger(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	event, err := models.GetEventById(database.DB, eventID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

// POST /events
func HandlePostEvents(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	event.UserID = userID

	if err := event.Save(database.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   event,
	})
}

// PUT /events/:id
func HandleUpdateEventById(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventID, err := utils.ParseToInteger(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	event.ID = eventID
	event.UserID = userID

	if err := event.Update(database.DB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated"})
}

// DELETE /events/:id
func HandleDeleteEventById(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventID, err := utils.ParseToInteger(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	if err := models.DeleteEvent(database.DB, eventID, userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
