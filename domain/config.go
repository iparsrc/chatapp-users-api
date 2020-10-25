package domain

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Connection string format: mongodb://[username:password@]host1[:port1]
	mongodbHost     = "mongodb_host"
	mongodbUsername = "mongodb_username"
	mongodbPassword = "mongodb_password"
)

var (
	db       *mongo.Database
	host     = os.Getenv(mongodbHost)
	username = os.Getenv(mongodbUsername)
	password = os.Getenv(mongodbPassword)
)

func init() {
	// 1. Set the default host if not provided.
	if host == "" {
		host = "localhost:27017"
	}

	// 2. Create connection string.
	ccStr := "mongodb://" + username + ":" + password + "@" + host
	if username == "" || password == "" {
		ccStr = "mongodb://" + host
	}

	// 3. Connect to the database.
	connectDB(ccStr)
}

func connectDB(URI string) {
	// 1. Prepare to connect the database.
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	// 2. Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// 3. Get the `users` database from the cluster.
	db = client.Database("users")
}
