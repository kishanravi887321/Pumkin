package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/dbs"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
	"github.com/kishanravi887321/pumkin/server-side/internal/routes"
)

func Run() error {
	cfg := config.Load()

	dbs.ConnectMongoDb()

	// Disable Gin debug and route registration logs.
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Keep panic recovery and your custom CORS middleware.
	router.Use(gin.Recovery(), middleware.CORS())

	// Don't trust forwarded IP headers from unspecified proxies.
	if err := router.SetTrustedProxies(nil); err != nil {
		return err
	}

	routes.Register(router, handlers.NewHealthHandler(cfg))
	log.Println("Server is listening at :", cfg.Port)

	return router.Run(fmt.Sprintf(":%s", cfg.Port))
}
