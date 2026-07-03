package router

import (
	"goportunities/handler"

	"github.com/gin-gonic/gin"
)

func InitializerRouter(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening/:id", handler.ShowByIdOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

}
