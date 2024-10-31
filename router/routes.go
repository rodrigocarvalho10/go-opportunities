package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/rodrigocarvalho10/go-opportunities/docs"
	"github.com/rodrigocarvalho10/go-opportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	//Initializer handler
	handler.InitializerHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.GetOpeningsHandler)
	}
}
