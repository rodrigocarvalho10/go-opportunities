package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	//Define a porta de inicialização do server
	router.Run(":8080")
}
