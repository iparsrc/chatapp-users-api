package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnectDB(URI string) {
	// Prepare to connect the database.
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	// db is a pointer to 'users' database.
	// db and its collections can be used inside the package.
	db = client.Database("users")
}
