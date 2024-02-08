package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	var collaborators []string
	if project.Collaborators != nil {
		collaborators = project.Collaborators
	}
	var description string
	if project.Description != nil {
		description = *project.Description
	} else {
		description = ""
	}
	return projects.Project{
		Name:          project.Name,
		Description:   description,
		Owner:         ownerId,
		Collaborators: handleCollaborators(collaborators, ownerId),
	}
}

func handleCollaborators(coll []string, owner primitive.ObjectID) []primitive.ObjectID {
	collIds := utils.MapTo(coll, func(collaborator string) primitive.ObjectID {
		id, err := primitive.ObjectIDFromHex(collaborator)
		if err != nil {
			return primitive.NilObjectID
		}
		return id
	})
	if !utils.Any(collIds, func(collaborator primitive.ObjectID) bool {
		return collaborator == owner
	}) {
		collIds = append(collIds, owner)
	}
	if len(collIds) > 1 {
		collIds[0], collIds[len(coll)-1] = collIds[len(coll)-1], collIds[0]
	}
	return collIds
}
