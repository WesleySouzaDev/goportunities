package handler

import (
	"fmt"
	"goportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpeningByIdHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, &id).Error; err != nil {
		logger.Errorf("error finding opening %v", err.Error())
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found with id %s", id))
		return
	}

	sendSuccess(ctx, "show-by-id-opening", opening)
}
