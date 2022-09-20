package routes

import (
	"net/http"

	"github.com/DustinDust/gin-blog-post/controllers"
	"github.com/DustinDust/gin-blog-post/middlewares"
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
		v1.GET("user/me", middlewares.JwtRequiredMiddleware(), controllers.UserController.FindMe)

		// RESOURCE /blog-post
		v1.GET("blog-post/:id", controllers.BlogPostController.FindById)
		v1.GET("blog-post", controllers.BlogPostController.FindAll)
		v1.POST("blog-post", middlewares.JwtRequiredMiddleware(), controllers.BlogPostController.Create)
		v1.PUT("blog-post/:id", middlewares.JwtRequiredMiddleware(), controllers.BlogPostController.Update)
		v1.DELETE("blog-post/:id", middlewares.JwtRequiredMiddleware(), controllers.BlogPostController.Delete)

		// RESOURCE /tag
		v1.GET("tag/:id", controllers.TagController.FindById)
		v1.GET("tag", controllers.TagController.FindAll)
		v1.POST("tag", middlewares.JwtRequiredMiddleware(), controllers.TagController.Create)
		v1.PUT("tag/:id", middlewares.JwtRequiredMiddleware(), controllers.TagController.Update)
		v1.DELETE("tag/:id", middlewares.JwtRequiredMiddleware(), controllers.TagController.Delete)

		// AUTH
		authorized := v1.Group("auth")
		{
			authorized.POST("test", middlewares.JwtRequiredMiddleware(), controllers.AuthController.JwtTest)
			authorized.POST("login", controllers.AuthController.Login)
			authorized.POST("register", controllers.AuthController.Register)
			authorized.POST("refresh", middlewares.RefreshTokenRequiredMiddleWare(), controllers.AuthController.Refresh)
		}
	}
}
