package opening

import (
	"fmt"
	"goportunities/handler"
	"goportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningByIdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func OpeningByIdHandler(ctx *gin.Context) {
	db, logger := handler.InitializeHandler()

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
