package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func userFromId(userId primitive.ObjectID) *users.User {
	user, err := app.App.Users.FindById(userId)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
func MapToQueryUser(user users.User) *model.User {
	return &model.User{
		ID:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Skills: utils.MapTo(user.Skills, func(skill util_types.Skill) string {
			return skill.String()
		}),
	}
}
func MapToUser(user model.NewUser) users.User {
	fmt.Println(user)
	var username string
	if user.Username != nil {
		username = *user.Username
	} else {
		username = utils.CreateUsername(user.FirstName, user.LastName)
	}
	//alternative
	// new_username := utils.CreateUsername(user.FirstName, user.LastName)
	// username := *utils.Ternary(user.Username != nil, user.Username, &new_username)
	fmt.Println(username)
	return users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  username,
		Email:     user.Email,
		Password:  user.Password,
		Skills: utils.MapTo(user.Skills, func(skill string) util_types.Skill {
			return util_types.Skill(skill)
		}),
		Projects: make([]projects.Project, 0),
	}
}
