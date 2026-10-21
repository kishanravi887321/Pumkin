package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
)

func Register(router *gin.Engine, health *handlers.HealthHandler) {
	router.GET("/health", health.Check)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
	})
}
