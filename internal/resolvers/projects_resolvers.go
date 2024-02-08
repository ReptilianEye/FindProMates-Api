package resolvers

import (
	"context"
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/auth"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProjectById(ctx context.Context, id string) (*projects.Project, error) {
	userId := auth.ForContext(ctx)
	if userId == "" {
		return nil, fmt.Errorf("access denied")
	}
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	projectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	project, err := app.App.Projects.FindById(projectId)
	if err != nil {
		return nil, err
	}
	if project.Owner != userIdObj {
		return nil, fmt.Errorf("access denied")
	}
	return project, nil
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
		SkillsNeeded: utils.MapTo(project.SkillsNeeded, func(skill util_types.Skill) string {
			return skill.String()
		}),
	}
}
func MapToProjectFromNew(project model.NewProject, ownerId primitive.ObjectID) projects.Project {
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
		Collaborators: handleCollaborators(collaborators, ownerId.Hex()),
	}
}
func UpdateProject(baseProject *projects.Project, project model.UpdatedProject) {
	baseProject.Name = utils.Elivis(project.Name, baseProject.Name)
	baseProject.Description = utils.Elivis(project.Description, baseProject.Description)
	if project.Collaborators != nil {
		baseProject.Collaborators = handleCollaborators(project.Collaborators, baseProject.Owner.Hex())
	}
	if project.SkillsNeeded != nil {
		skills := handleSkillsNeeded(project.SkillsNeeded)
		fmt.Println(skills)
		baseProject.SkillsNeeded = skills
	}
}

func handleCollaborators(collabs []string, owner string, base ...primitive.ObjectID) []primitive.ObjectID {
	safePrimitive := func(a string) primitive.ObjectID {
		id, err := primitive.ObjectIDFromHex(a)
		if err != nil {
			log.Fatal()
		}
		return id
	}
	return utils.MergeSlices(base, collabs, primitive.ObjectID.Hex, safePrimitive, owner)
}
func handleSkillsNeeded(skills []string, base ...util_types.Skill) []util_types.Skill {
	toString := func(s util_types.Skill) string {
		return s.String()
	}
	fromString := (func(s string) util_types.Skill {
		skill := util_types.Skill(s)
		if !skill.IsValid() {
			log.Fatal()
		}
		return skill
	})
	return utils.MergeSlices(base, skills, toString, fromString)
}
