package controllers

import (
	"log"
	"strconv"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/gin-gonic/gin"
)

func PostBlog(c *gin.Context) {
	var newBlog models.CreateBlog

	if err := c.BindJSON(&newBlog); err != nil {
		return
	}

	blog, err := models.CreateBlogs(newBlog)

	if err != nil {
		c.JSON(404, gin.H{"error": "Blog Couldn't Be Created", "message": err})
	}

	c.JSON(200, gin.H{"data": blog})
}

func ReadBlogs(c *gin.Context) {
	blogs, err := models.GetBlogs()
	checkErr(err)

	if blogs == nil {
		c.JSON(404, gin.H{"error": "No Blogs Found"})
	} else {
		c.JSON(200, gin.H{"data": blogs})
	}
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	checkErr(err)

	blog, err := models.GetBlogById(convertedId)

	checkErr(err)

	if blog == nil {
		c.JSON(404, gin.H{"error": "No Blog Found By Given Id"})
	} else {
		c.JSON(200, gin.H{"data": blog})
	}
}

func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	checkErr(err)

	blog, err := models.GetBlogById(convertedId)

	checkErr(err)

	if err := c.BindJSON(&blog); err != nil {
		return
	}

	updatedBlog, err := models.UpdateBlogById(*blog)

	if err != nil {
		c.JSON(404, gin.H{"error": "Blog Couldn't Be Updated", "message": err})
	}

	c.JSON(200, gin.H{"message": "Blog Updated Successfully", "data": updatedBlog})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}