package main

import (
	"log"
	"strconv"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/gin-gonic/gin"
)

// TODO : Move all the Routing Logic Into Router files
// TODO : Move all the Business Logic into service files
// TODO : Create User Service Endpoints

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

    r := gin.Default()
	blogRouter := r.Group("/blogs")
	{
		blogRouter.GET("/", readBlogs)
		blogRouter.GET("/:id",getBlogById)
		blogRouter.POST("/",postBlog)
		blogRouter.PUT("/:id",updateBlog)
	}
    r.Run("localhost:8080")
}

func postBlog(c *gin.Context){
	var newBlog models.CreateBlog

	if err := c.BindJSON(&newBlog); err != nil{
		return 
	}

	blog,err := models.CreateBlogs(newBlog)

	if err != nil {
		c.JSON(404,gin.H{"error":"Blog Couldn't Be Created","message":err})
	}

	c.JSON(200,gin.H{"data":blog})
}

func readBlogs(c *gin.Context){
	blogs,err := models.GetBlogs()
	checkErr(err)

	if blogs == nil {
		c.JSON(404,gin.H{"error":"No Blogs Found"})
	}else{
		c.JSON(200,gin.H{"data":blogs})
	}
}

func getBlogById(c *gin.Context){
	id := c.Param("id")
	convertedId,err := strconv.Atoi(id)
	checkErr(err)

	blog,err := models.GetBlogById(convertedId)

	checkErr(err)
	
	if blog == nil {
		c.JSON(404,gin.H{"error":"No Blog Found By Given Id"})
	}else{
		c.JSON(200,gin.H{"data":blog})
	}
}

func updateBlog(c *gin.Context){
	id := c.Param("id")
	convertedId,err := strconv.Atoi(id)
	checkErr(err)

	blog,err := models.GetBlogById(convertedId)

	checkErr(err)

	if err := c.BindJSON(&blog); err != nil{
		return 
	}

	updatedBlog,err := models.UpdateBlogById(*blog)

	if err != nil {
		c.JSON(404,gin.H{"error":"Blog Couldn't Be Updated","message":err})
	}

	c.JSON(200,gin.H{"message":"Blog Updated Successfully","data":updatedBlog})
}

func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}
