package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var headerKeyJson = "Content-Type"
var headerValueJson = "application/json"

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header(headerKeyJson, headerValueJson)
	ctx.JSON(code, gin.H{
		"message": message,
		"code":    code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data any) {
	ctx.Header(headerKeyJson, headerValueJson)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successful", op),
		"data":    data,
	})
}
