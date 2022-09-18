package routes

import (
	"net/http"

	"github.com/DustinDust/gin-blog-post/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouteV1(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello world",
			})
		})

		// RESOURCE /user
		v1.GET("user/:id", controllers.UserController.FindById)
		v1.POST("user", controllers.UserController.Create)
		v1.PUT("user/:id", controllers.UserController.Update)
		v1.DELETE("user/:id", controllers.UserController.Delete)

		// RESOURCE /blog-post
		v1.GET("blog-post/:id", controllers.BlogPostController.FindById)
		v1.GET("blog-post", controllers.BlogPostController.FindAll)
		v1.POST("blog-post", controllers.BlogPostController.Create)
		v1.PUT("blog-post/:id", controllers.BlogPostController.Update)
		v1.DELETE("blog-post/:id", controllers.BlogPostController.Delete)

		// RESOURCE /tag
		v1.GET("tag/:id", controllers.TagController.FindById)
		v1.GET("tag", controllers.TagController.FindAll)
		v1.POST("tag", controllers.TagController.Create)
		v1.PUT("tag/:id", controllers.TagController.Update)
		v1.DELETE("tag/:id", controllers.TagController.Delete)
	}
}
