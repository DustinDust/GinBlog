package main

import (
	"log"
	"net/http"

	"github.com/DustinDust/gin-blog-post/db"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		log.Printf("Error: %v", err)
		return
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	err := r.Run(":3000")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
