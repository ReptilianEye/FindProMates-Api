package main

import (
	"context"
	"example/FindProMates-Api/internal/models/mongodb"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Projects *mongodb.ProjectModel
	Users    *mongodb.UserModel
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		log.Fatal("DB_NAME is not set")
	}
	// db := client.Database(db_name)
	// projectsColl := db.Collection(dbschema.ProjectCollection)

}
