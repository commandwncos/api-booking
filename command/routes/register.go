package routes

import (
	"net/http"
	"strconv"

	"github.com/commandwncos/api-booking/command/private/database"
	"github.com/commandwncos/api-booking/command/utils"
	"github.com/commandwncos/api-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	// impede o dono do evento de se registrar
	isOwner, err := models.IsEventOwner(database.DB, eventID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	if isOwner {
		c.JSON(http.StatusBadRequest, gin.H{"message": "event owner cannot register"})
		return
	}

	err = models.RegisterUserForEvent(database.DB, userID, eventID)
	if err != nil {
		if err == models.ErrAlreadyRegistered {
			c.JSON(http.StatusConflict, gin.H{"message": "already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func cancelForEvent(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	rows, err := models.CancelRegistration(database.DB, userID, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "registration not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration canceled"})
}
