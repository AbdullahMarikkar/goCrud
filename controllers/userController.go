package controllers

import (
	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {
	var newUser models.CreateUserDto

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user, err := models.CreateUser(newUser)

	if err != nil {
		c.JSON(400, gin.H{"error": "User Couldn't Be Created", "message": err})
	}

	c.JSON(200, gin.H{"data": user})
}