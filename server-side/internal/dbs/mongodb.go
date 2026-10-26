package dbs

import (
	"context"
	"fmt"
	"time"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDb() (*mongo.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	url := cfg.DatabaseURL

	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	fmt.Println("MongoDB connected successfully!")
	return client, nil
}
