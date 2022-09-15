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
		v1.GET("user/:id", controllers.UserController.FindById)
		v1.POST("user", controllers.UserController.Create)
		v1.PUT("user/:id", controllers.UserController.Update)
		v1.DELETE("user/:id", controllers.UserController.Delete)
	}
}
