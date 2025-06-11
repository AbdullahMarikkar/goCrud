package controllers

import (
	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/AbdullahMarikkar/goCrud/utils"
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

func LogInUserController(c *gin.Context){
	type LogInInput struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var logInInput LogInInput

	if err := c.BindJSON(&logInInput); err != nil {
		return
	}

	user , err := models.AuthorizeUser(logInInput.Email,logInInput.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Email or Password", "message": err})
	}

	token,err := utils.CreateToken(logInInput.Email)

	if err != nil {
		c.JSON(400, gin.H{"error": "Create Token Error", "message": err})
	}

	c.SetCookie("accessToken",token,1000*60*60*24,"/","localhost",false,true)
	c.JSON(200,gin.H{"data":user})
}