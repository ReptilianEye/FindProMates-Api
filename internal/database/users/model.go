package users

import (
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/util_types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollection string = "users"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName,omitempty" `
	LastName  string             `bson:"lastName,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Skills    []util_types.Skill `bson:"skills,omitempty"`
	Projects  []projects.Project `bson:"projects,omitempty"`
}
