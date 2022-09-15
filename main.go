package main

import (
	"log"

	"github.com/DustinDust/gin-blog-post/controllers"
	"github.com/DustinDust/gin-blog-post/db"
	"github.com/DustinDust/gin-blog-post/models"
	"github.com/DustinDust/gin-blog-post/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := models.InitModel(); err != nil {
		log.Printf("Error: %v", err)
	}

	controllers.InitController()

	r := gin.Default()
	routes.InitRouteV1(r)

	err := r.Run(":3000")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
