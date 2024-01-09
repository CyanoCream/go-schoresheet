package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var mongoDB *mongo.Database

func InitMongoDB() *mongo.Database {
	// Mengambil nilai-nilai konfigurasi dari variabel lingkungan
	dbHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_NAME")
	dbPort := os.Getenv("MONGO_PORT")
	dbUsername := os.Getenv("MONGO_USERNAME")
	dbPassword := os.Getenv("MONGO_PASSWORD")

	// Membuat URI koneksi MongoDB
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Set up options for the MongoDB client
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Access a specific database
	mongoDB = client.Database("mydb") // Ganti "mydb" dengan nama database yang sesuai

	return mongoDB
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}
