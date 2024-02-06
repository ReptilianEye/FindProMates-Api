package models

import (
	"example/FindProMates-Api/internal/models/util_types"
)

const UserCollection string = "users"

type User struct {
	FirstName string             `bson:"first_name,omitempty" `
	LastName  string             `bson:"last_name,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Skills    []util_types.Skill `bson:"interests"`
}
