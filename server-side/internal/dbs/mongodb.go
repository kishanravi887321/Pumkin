package dbs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDb() {

	url, err := config.Dburl()

	client, err := mongo.Connect(options.Client().ApplyURI(url))

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {

		log.Fatal("MongoDB connection failed:", err)

	}

	fmt.Println("MongoDB connected successfully!")
}
