package handler

import (
	"fmt"
	"goportunities/schemas"
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

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
type ShowOpeningByIdResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
