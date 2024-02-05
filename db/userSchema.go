package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"id"`
	FirstName string             `bson:"first_name,omitempty" `
	LastName  string             `bson:"last_name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Skills    []string           `bson:"interests,omitempty"`
}
