package users

import (
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/util_types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollection string = "users"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name,omitempty" `
	LastName  string             `bson:"last_name,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Skills    []util_types.Skill `bson:"interests"`
	Projects  []projects.Project `bson:"projects"`
}
