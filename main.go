package main

import (
	"log"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/gin-gonic/gin"
)



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
	c.JSON(200,gin.H{"message":"A New Blog Created"})
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
	c.JSON(200,gin.H{"message":"Blog of ID : "})
}

func updateBlog(c *gin.Context){
	c.JSON(200,gin.H{"message":"Blog Updated Successfully"})
}

func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}
