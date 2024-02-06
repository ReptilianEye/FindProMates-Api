package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readFromDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	db := client.Database(os.Getenv("DB_NAME"))
	usersColl := db.Collection("users")
	projectsColl := db.Collection("projects")
	file, err := os.ReadFile("dbschema/db.json")
	if err != nil {
		log.Fatal(err)
	}
	var db_sample map[string]interface{}
	err = json.Unmarshal(file, &db_sample)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range db_sample["users"].([]interface{}) {
		_, err := usersColl.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
		}
	}
	for _, project := range db_sample["projects"].([]interface{}) {

		res, err := projectsColl.InsertOne(context.Background(), project)
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
		} else {
			fmt.Println(res.InsertedID)
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	readFromDB()

}
