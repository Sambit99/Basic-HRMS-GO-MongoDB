package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Collection

func ConnectDB() {
	if err := godotenv.Load("../../.env.development"); err != nil {
		fmt.Println("Error loading .env file")
	}

	MONGO_CONNECTION_STRING := os.Getenv("MONGO_CONNECTION_STRING")
	MONGO_DATABASE_NAME := os.Getenv("MONGO_DATABASE_NAME")
	MONGO_COLLECTION_NAME := os.Getenv("MONGO_COLLECTION_NAME")

	opts := options.Client().ApplyURI(MONGO_CONNECTION_STRING)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	collection := client.Database(MONGO_DATABASE_NAME).Collection(MONGO_COLLECTION_NAME)
	db = collection
}

func GetDB() *mongo.Collection {
	return db
}
