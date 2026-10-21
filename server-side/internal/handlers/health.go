package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

type HealthHandler struct {
	config config.Config
}

func NewHealthHandler(cfg config.Config) *HealthHandler {
	return &HealthHandler{config: cfg}
}

func (h *HealthHandler) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"environment": h.config.Environment,
	})
}
