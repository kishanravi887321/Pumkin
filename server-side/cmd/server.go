package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/kishanravi887321/pumkin/server-side/internal/app"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/dbs"
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

	router, err := app.NewRouter(cfg, mongoClient)

	if err != nil {
		return err
	}

	log.Println("Server is listening at :", cfg.Port)

	return router.Run(fmt.Sprintf(":%s", cfg.Port))
}
