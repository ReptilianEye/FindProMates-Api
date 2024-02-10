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
	"slices"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProjectByStrId(id string) (*projects.Project, error) {
	projectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.Projects.FindById(projectId)
}
func ProjectByObjId(id primitive.ObjectID) *projects.Project {
	project, err := app.App.Projects.FindById(id)
	if err != nil {
		log.Panic(err)
	}
	return project
}
func CanQueryProject(project *projects.Project, user *users.User) bool {
	return project.Public || CanMutateProject(project, user)
}

// CanMutateProject checks if the user is the owner or a collaborator of the project
func CanMutateProject(project *projects.Project, user *users.User) bool {
	return IsOwner(project, user) || slices.Contains(project.Collaborators, user.ID)
}
func IsOwner(project *projects.Project, user *users.User) bool {
	return project.Owner == user.ID
}
func ProjectsOwnedByUser(user *users.User) ([]*model.Project, error) {
	ownedProjects, err := app.App.Projects.FindByOwner(user.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(ownedProjects, MapToQueryProject), nil
}
func ProjectsCollaboratedByUser(user *users.User) ([]*model.Project, error) {
	collaboratedProjects, err := app.App.Projects.FindAllUserIsCollaborator(user.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(collaboratedProjects, MapToQueryProject), nil
}

func MapToQueryProject(project *projects.Project) *model.Project {
	return &model.Project{
		ID:               project.ID.Hex(),
		Name:             project.Name,
		Description:      project.Description,
		Owner:            MapToQueryUser(UserByObjId(project.Owner)),
		Public:           project.Public,
		CompletionStatus: project.CompletionStatus.String(),
		Collaborators: utils.MapTo(project.Collaborators, func(collaborator primitive.ObjectID) *model.User {
			return MapToQueryUser(UserByObjId(collaborator))
		}),
		SkillsNeeded: utils.MapTo(project.SkillsNeeded, func(skill util_types.Skill) string {
			return skill.String()
		}),
	}
}
func MapToProjectFromNew(newProject model.NewProject, ownerId primitive.ObjectID) (*projects.Project, error) {
	collaborators := utils.Elivis(&newProject.Collaborators, []string{})
	collaboratorsObj, err := handleCollaborators(collaborators, ownerId.Hex())
	if err != nil {
		return &projects.Project{}, err
	}
	skillsNeeded := utils.Elivis(&newProject.SkillsNeeded, []string{})
	skillsNeededObj, err := handleSkillsNeeded(skillsNeeded)
	if err != nil {
		return &projects.Project{}, err
	}
	return &projects.Project{
		Owner:            ownerId,
		Name:             newProject.Name,
		Description:      utils.Elivis(newProject.Description, ""),
		Public:           utils.Elivis(newProject.Public, false),
		CompletionStatus: util_types.InProgress,
		Collaborators:    collaboratorsObj,
		SkillsNeeded:     skillsNeededObj,
	}, nil
}
func UpdateProject(baseProject *projects.Project, updatedProject model.UpdatedProject) error {
	baseProject.Name = utils.Elivis(updatedProject.Name, baseProject.Name)
	baseProject.Description = utils.Elivis(updatedProject.Description, baseProject.Description)
	baseProject.Public = utils.Elivis(updatedProject.Public, baseProject.Public)
	if updatedProject.CompletionStatus != nil {
		status := util_types.CompletionStatus(*updatedProject.CompletionStatus)
		if !status.IsValid() {
			return fmt.Errorf("invalid completion status")
		}
		baseProject.CompletionStatus = status
	}

	if updatedProject.Collaborators != nil {
		collabs, err := handleCollaborators(updatedProject.Collaborators, baseProject.Owner.Hex())
		if err != nil {
			return err
		}
		baseProject.Collaborators = collabs
	}
	if updatedProject.SkillsNeeded != nil {
		skills, err := handleSkillsNeeded(updatedProject.SkillsNeeded, baseProject.SkillsNeeded...)
		if err != nil {
			return err
		}
		baseProject.SkillsNeeded = skills
	}

	return nil
}
func PublicProjects() ([]*model.Project, error) {
	projectsArr, err := app.App.Projects.All()
	if err != nil {
		return nil, err
	}
	publicProjectsArr := utils.Filter(projectsArr, func(p *projects.Project) bool { return p.Public })
	return utils.MapTo(publicProjectsArr, MapToQueryProject), nil
}
func handleCollaborators(collabs []string, owner string, base ...primitive.ObjectID) ([]primitive.ObjectID, error) {
	toString := func(id primitive.ObjectID) (string, error) {
		return id.Hex(), nil
	}
	safePrimitive := func(strId string) (primitive.ObjectID, error) {
		id, err := primitive.ObjectIDFromHex(strId)
		if err != nil {
			return primitive.ObjectID{}, fmt.Errorf("provided collaborator id: '%s' is invalid: %v", strId, err)
		}
		return id, nil

	}
	return utils.MergeSlices(base, collabs, toString, safePrimitive, owner)
}
func handleSkillsNeeded(skills []string, base ...util_types.Skill) ([]util_types.Skill, error) {
	toString := func(s util_types.Skill) (string, error) {
		return s.String(), nil
	}
	fromString := func(s string) (util_types.Skill, error) {
		skill := util_types.Skill(s)
		if !skill.IsValid() {
			return skill, fmt.Errorf("invalid skill")
		}
		return skill, nil

	}
	return utils.MergeSlices(base, skills, toString, fromString)
}
