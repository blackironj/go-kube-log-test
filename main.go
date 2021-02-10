package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		log.Info("Ping!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			log.Error("wrong query param")

			c.JSON(400, gin.H{
				"message": "Need a name query param",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello! " + name,
		})
	})

	r.Run(":8080")
}
