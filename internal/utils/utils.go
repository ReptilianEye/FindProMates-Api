package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MapTo[E, V comparable](arr []E, mapper func(E) V) []V {
	mapped := make([]V, len(arr))
	for i, v := range arr {
		mapped[i] = mapper(v)
	}
	return mapped
}

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
