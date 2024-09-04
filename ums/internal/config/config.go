package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
	mongoURI := os.Getenv("DB_URL")
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client
}

func GetUserCollection(client *mongo.Client) *mongo.Collection {
	name := os.Getenv("DB_NAME")
	return client.Database(name).Collection("users")
}
