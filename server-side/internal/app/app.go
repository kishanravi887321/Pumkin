package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
	"github.com/kishanravi887321/pumkin/server-side/internal/repository"
	"github.com/kishanravi887321/pumkin/server-side/internal/routes"
	"github.com/kishanravi887321/pumkin/server-side/internal/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(cfg config.Config, mongoClient *mongo.Client) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery(), middleware.CORS())

	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("configure trusted proxies: %w", err)
	}

	userCollection := mongoClient.Database(cfg.MongoDatabaseName).Collection("users")
	userRepository := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	routes.UserRouter(router, userHandler)

	return router, nil
}
