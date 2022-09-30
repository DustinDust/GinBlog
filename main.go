package main

import (
	"log"

	"github.com/DustinDust/gin-blog-post/controllers"
	"github.com/DustinDust/gin-blog-post/db"
	_ "github.com/DustinDust/gin-blog-post/docs"
	"github.com/DustinDust/gin-blog-post/middlewares"
	"github.com/DustinDust/gin-blog-post/models"
	"github.com/DustinDust/gin-blog-post/routes"
	"github.com/DustinDust/gin-blog-post/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Blogpost API
// @version 1.0.0
// @description Blogpost API written in Golang with Gin & Gorm
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @Basepath /v1
// @host localhost:9090

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := db.InitDB(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	middlewares.Init()
	services.Init()
	if err := models.InitModel(); err != nil {
		log.Printf("Error: %v", err)
	}

	controllers.InitController()

	r := gin.Default()
	routes.InitRouteV1(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":9090")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
