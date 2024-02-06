package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const ProjectCollection string = "projects"

type Project struct {
	Owner         primitive.ObjectID   `bson:"owner,omitempty"`
	Name          string               `bson:"name,omitempty"`
	Description   string               `bson:"description,omitempty"`
	Collaborators []primitive.ObjectID `bson:"collaborators,omitempty"`
}
