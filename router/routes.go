package router

import (
	"goportunities/handler"
	"goportunities/handler/opening"

	"github.com/gin-gonic/gin"

	docs "goportunities/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializerRouter(router *gin.Engine) {

	basePath := "/api/v1"
	// Initialize handler
	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// Openings
		v1.GET("/openings", opening.ListOpeningsHandler)
		v1.GET("/opening", opening.OpeningByIdHandler)
		v1.POST("/opening", opening.CreateOpeningHandler)
		v1.PUT("/opening", opening.UpdateOpeningHandler)
		v1.DELETE("/opening", opening.DeleteOpeningHandler)

	}

	// Initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
