package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/dbs"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
	"github.com/kishanravi887321/pumkin/server-side/internal/routes"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	fmt.Println("Server is starting in", cfg.Environment, "mode...")

	mongoClient, err := dbs.ConnectMongoDb()
	if err != nil {
		return err
	}
	defer func() {
		_ = mongoClient.Disconnect(context.Background())
	}()

	postgresDB, err := dbs.ConnectPostgres()
	if err != nil {
		return err
	}
	defer postgresDB.Close()

	// Disable Gin debug and route registration logs.
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())

	// Keep panic recovery and your custom CORS middleware.
	router.Use(gin.Recovery(), middleware.CORS())

	// Don't trust forwarded IP headers from unspecified proxies.
	if err := router.SetTrustedProxies(nil); err != nil {
		return err
	}

	routes.UserRouter(router)

	log.Println("Server is listening at :", cfg.Port)

	return router.Run(fmt.Sprintf(":%s", cfg.Port))
}
