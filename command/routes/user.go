package routes

import (
	"fmt"
	"net/http"

	"github.com/commandwncos/api-booking/command/private/database"
	"github.com/commandwncos/api-booking/command/utils"
	"github.com/commandwncos/api-booking/models"
	g "github.com/gin-gonic/gin"
)

func signup(c *g.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, g.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save(database.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, g.H{"message": "Could not save user"})
		return

	}

	c.JSON(http.StatusCreated, g.H{"message": "User created susccessfully."})
}

func login(c *g.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, g.H{"message": "Could not parse request data"})
		return
	}

	valid, err := user.ValidateCredentials(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, g.H{"message": "Could not authenticate user"})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, g.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateJsonWebToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, g.H{"message": "Could not authenticate user"})
		return
	}
	c.JSON(http.StatusOK, g.H{"message": "Login successfully", "token": token})
}
