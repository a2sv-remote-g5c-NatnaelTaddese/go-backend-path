package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
)

func InitDB(mongoURI string) error {
	var err error
	ctx = context.Background()

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")
	collection = client.Database("taskmanager").Collection("tasks")

	return nil
}

func CloseDB() {
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
