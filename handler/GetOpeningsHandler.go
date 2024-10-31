package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-opportunities/schemas"
)

func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
