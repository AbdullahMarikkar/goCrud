package routers

import (
	"github.com/AbdullahMarikkar/goCrud/controllers"
	"github.com/AbdullahMarikkar/goCrud/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run("localhost:7080")
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	blogRouter := router.Group("/blogs",middleware.AuthorizeMiddleware)
	userRouter := router.Group("/users")
	// blogRouter.Use(middleware)
	{
		blogRouter.GET("/",controllers.ReadBlogs )
		blogRouter.GET("/:id",controllers.GetBlogById)
		blogRouter.POST("/",controllers.PostBlog)
		blogRouter.PUT("/:id",controllers.UpdateBlog)
	}
	{
		userRouter.POST("/",controllers.CreateUserController)
		userRouter.POST("/login",controllers.LogInUserController)
	}
	return router
}