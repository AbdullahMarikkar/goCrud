package middleware

import (
	"github.com/AbdullahMarikkar/goCrud/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware(c *gin.Context){

	cookie,err := c.Request.Cookie("accessToken")

	if err != nil {
		c.JSON(401, gin.H{"error": "Couldn't Find Access Token", "message": err})
	}

	valid,err := utils.VerifyToken(cookie.Value)

	if err != nil ||  !valid{
		c.JSON(401, gin.H{"error": "Invalid Token", "message": err})
	}

	c.Next()
}