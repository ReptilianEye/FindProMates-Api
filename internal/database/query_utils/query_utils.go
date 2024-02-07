package queryutils

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
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
	username := utils.Ternary(user.Username != nil, *user.Username, createUsername(user))
	return users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  username,
		Email:     user.Email,
		Password:  user.Password,
		Skills: utils.MapTo(user.Skills, func(skill string) util_types.Skill {
			return util_types.Skill(skill)
		}),
	}
}
func MapToQueryProject(project projects.Project) *model.Project {
	return &model.Project{
		ID:          project.ID.Hex(),
		Name:        project.Name,
		Description: project.Description,
		Owner:       MapToQueryUser(*userFromId(project.Owner)),
		Collaborators: utils.MapTo(project.Collaborators, func(collaborator primitive.ObjectID) *model.User {
			return MapToQueryUser(*userFromId(collaborator))
		}),
	}
}
func MapToProject(project model.NewProject, ownerId primitive.ObjectID) projects.Project {
	collaborators := utils.MapTo(project.Collaborators, func(collaborator string) primitive.ObjectID {
		return primitive.ObjectID{}
	})
	if !utils.Any(collaborators, func(collaborator primitive.ObjectID) bool {
		return collaborator == ownerId
	}) {
		collaborators = append(collaborators, ownerId)
	}
	return projects.Project{
		Name:          project.Name,
		Description:   utils.Ternary(project.Description != nil, *project.Description, ""),
		Owner:         ownerId,
		Collaborators: collaborators,
	}
}

func createUsername(user model.NewUser) string {
	return user.FirstName + "_" + user.LastName
}
