package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Db *mongo.Database

func InitDB() context.CancelFunc {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	Db = client.Database(os.Getenv("DB_NAME"))
	return cancel
}
func CloseDB() {
	client.Disconnect(context.Background())
}
