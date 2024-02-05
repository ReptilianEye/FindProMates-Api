package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID            primitive.ObjectID   `bson:"id"`
	Owner         primitive.ObjectID   `bson:"owner,omitempty"`
	Name          string               `bson:"name,omitempty"`
	Description   string               `bson:"description,omitempty"`
	Collaborators []primitive.ObjectID `bson:"collaborators,omitempty"`
}
