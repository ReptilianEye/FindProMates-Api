package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"log"
	"slices"
	"sort"

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
		baseProject.CompletionStatus = util_types.CompletionStatus(*updatedProject.CompletionStatus)
		if err := baseProject.CompletionStatus.IsValid(); err != nil {
			return err
		}
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
	publicProjects, err := app.App.Projects.AllPublic()
	if err != nil {
		return nil, err
	}
	return utils.MapTo(publicProjects, MapToQueryProject), nil
}
func RecommendedProjects(user *users.User) ([]*model.Project, error) {
	usersSkills, err := utils.ToSet(user.Skills, func(s util_types.Skill) (string, error) { return s.String(), nil })
	if err != nil {
		return nil, err
	}
	projects, err := app.App.Projects.AllPublic()
	if err != nil {
		return nil, err
	}
	buckets := map[int][]*model.Project{}
	for _, project := range projects {
		matchCount := 0
		for _, skill := range project.SkillsNeeded {
			if usersSkills.Contains(skill.String()) {
				matchCount++
			}
		}
		if matchCount > 0 {
			buckets[matchCount] = append(buckets[matchCount], MapToQueryProject(project))
		}
	}
	keys := utils.Keys(buckets)
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	result := []*model.Project{}
	for _, key := range keys {
		result = append(result, buckets[key]...)
	}
	return result, nil
}

func handleCollaborators(collabs []string, owner string, base ...primitive.ObjectID) ([]primitive.ObjectID, error) {
	return utils.MergeSlices(base, collabs, utils.SafeIDToString, utils.SafeStringToID, owner)
}
func handleSkillsNeeded(skills []string, base ...util_types.Skill) ([]util_types.Skill, error) {
	toString := func(s util_types.Skill) (string, error) {
		return s.String(), nil
	}
	fromString := func(s string) (util_types.Skill, error) {
		skill := util_types.Skill(s)
		if err := skill.IsValid(); err != nil {
			return "", err
		}
		return skill, nil

	}
	return utils.MergeSlices(base, skills, toString, fromString)
}
