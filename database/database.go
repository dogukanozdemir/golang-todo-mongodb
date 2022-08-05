package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = CreateMongoClient()

func CreateMongoClient() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file!")
	}
	MongoDbURI := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDbURI))
	if err != nil {
		log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	fmt.Println("Connected to MONGO!")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-mongodb").Collection(collectionName)
}
