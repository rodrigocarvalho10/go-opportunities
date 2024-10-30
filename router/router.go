package router

import (
	"github.com/gin-gonic/gin"

)

func Initialize() {

	router *gin.Engine := gin.Default()

	router.GET(relativePath: "/ping", handlers ... : func (c *gin.Context)  {
		c.JSON(code: 200, obj: gin.H{
			"message": "pong",
		})
		
	})

	router.Run(addr ... : ":8080")
}
