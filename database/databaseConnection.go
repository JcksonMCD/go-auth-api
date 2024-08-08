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

// Annotated for future reference as firt time using MongoDB!

// loads environment variables, creates a client, and connects to MongoDB.
func DBinstance() (*mongo.Client, error) {
	// Load environment variables from the .env file.
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Retrieve the MongoDB connection string from environment variables.
	mongoURI := os.Getenv("MONGODB_URL")

	// Create a new MongoDB client options object.
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create a new MongoDB client using the options object.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %v", err)
	}

	// Create a context with a timeout of 10 seconds for connection verification.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure that the context is cancelled to free up resources.

	// Verify the MongoDB connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	return client, nil
}

// Global variable to hold the MongoDB client instance.
var Client *mongo.Client

// init function is executed automatically when the program starts.
// It initializes the MongoDB client and handles any initialization errors.
func init() {
	// Call the DBinstance function to create and connect the MongoDB client.
	var err error
	Client, err = DBinstance()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}
}

// OpenCollection returns a specific collection from the MongoDB database.
// It uses the client to access the specified collection within the database.
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Access the database name from environment variables and open the specified collection.
	return client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collectionName)
}
