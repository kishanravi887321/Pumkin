package dbs

import (
	"context"
	"fmt"
	"time"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDb() {

	url := config.Dburl()

	client, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {

		fmt.Println("some error occured durign Mongo Connection")

	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("MongoDB connection failed:", err)
		return
	}

	fmt.Println("MongoDB connected successfully!")
}
