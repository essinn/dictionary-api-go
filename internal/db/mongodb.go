package db

import (
	"context"
	"dictionary-api/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToMongoDB(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return fmt.Errorf("Failed to connect to MongoDB %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to ping MongoDB %v", err)
	}

	Client = client
	return nil
}

func GetCollection(cfg *config.Config) *mongo.Collection {
	return Client.Database(cfg.DBName).Collection(cfg.Collection)
}
