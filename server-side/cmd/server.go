package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
	"github.com/kishanravi887321/pumkin/server-side/internal/routes"
)

func Run() error {
	cfg := config.Load()
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	routes.Register(router, handlers.NewHealthHandler(cfg))
	return router.Run(fmt.Sprintf(":%s", cfg.Port))
}
