package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svvictorelias/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 201 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, http.StatusOK, "listOpenings", openings)
}
