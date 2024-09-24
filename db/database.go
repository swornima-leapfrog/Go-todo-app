package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	TodoDetails *mongo.Collection
) 


func init() {
	err:= godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	hostname := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, hostname, port)

	clientOptions := options.Client().ApplyURI(uri)

	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = MongoClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Could not ping mongo db", err)
	}

	fmt.Println("MongoDb connected successfully!!")
	// // Set the collection to use for TODO items
	TodoDetails = MongoClient.Database("todo").Collection("todos")

}
