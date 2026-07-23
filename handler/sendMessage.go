package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var headerKeyJson = "Content-Type"
var headerValueJson = "application/json"

func SendError(ctx *gin.Context, code int, message string) {
	ctx.Header(headerKeyJson, headerValueJson)
	ctx.JSON(code, gin.H{
		"message": message,
		"code":    code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data any) {
	ctx.Header(headerKeyJson, headerValueJson)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successful", op),
		"data":    data,
	})
}

func ErrParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required.", param, typ)
}
