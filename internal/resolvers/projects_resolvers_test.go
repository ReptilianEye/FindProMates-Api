package resolvers

// import (
// 	"example/FindProMates-Api/graph/model"
// 	"example/FindProMates-Api/internal/database/projects"
// 	"example/FindProMates-Api/internal/database/util_types"
// 	"example/FindProMates-Api/internal/pkg/utils"
// 	"testing"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // does not work because cannot connect to database
// func TestMapToQueryProject(t *testing.T) {
// 	id := primitive.NewObjectID()
// 	name := "Test Project"
// 	description := "Test Description"
// 	owner := primitive.NewObjectID()
// 	collaborators := []primitive.ObjectID{
// 		primitive.NewObjectID(),
// 		primitive.NewObjectID(),
// 	}
// 	skillsNeeded := []util_types.Skill{
// 		util_types.Skill(util_types.Go),
// 		util_types.Skill(util_types.Java),
// 	}

// 	project := projects.Project{
// 		ID:            id,
// 		Name:          name,
// 		Description:   description,
// 		Owner:         owner,
// 		Collaborators: collaborators,
// 		SkillsNeeded:  skillsNeeded,
// 	}

// 	expectedProject := &model.Project{
// 		ID:          id.Hex(),
// 		Name:        name,
// 		Description: description,
// 		Owner:       &model.User{ID: owner.Hex()},
// 		Collaborators: utils.MapTo(project.Collaborators, func(collaborator primitive.ObjectID) *model.User {
// 			return &model.User{ID: collaborator.Hex()}
// 		}),
// 		SkillsNeeded: []string{util_types.Go.String(), util_types.Java.String()},
// 	}

// 	result := MapToQueryProject(project)

// 	if result.ID != expectedProject.ID {
// 		t.Errorf("MapToQueryProject() ID = %s, want %s", result.ID, expectedProject.ID)
// 	}

// 	if result.Name != expectedProject.Name {
// 		t.Errorf("MapToQueryProject() Name = %s, want %s", result.Name, expectedProject.Name)
// 	}

// 	if result.Description != expectedProject.Description {
// 		t.Errorf("MapToQueryProject() Description = %s, want %s", result.Description, expectedProject.Description)
// 	}

// 	if result.Owner.ID != expectedProject.Owner.ID {
// 		t.Errorf("MapToQueryProject() Owner = %s, want %s", result.Owner.ID, expectedProject.Owner.ID)
// 	}

// 	if len(result.Collaborators) != len(expectedProject.Collaborators) {
// 		t.Errorf("MapToQueryProject() Collaborators = %v, want %v", result.Collaborators, expectedProject.Collaborators)
// 	}

// 	for i, v := range result.Collaborators {
// 		if v.ID != expectedProject.Collaborators[i].ID {
// 			t.Errorf("MapToQueryProject() Collaborators = %v, want %v", result.Collaborators, expectedProject.Collaborators)
// 		}
// 	}

// 	if len(result.SkillsNeeded) != len(expectedProject.SkillsNeeded) {
// 		t.Errorf("MapToQueryProject() SkillsNeeded = %v, want %v", result.SkillsNeeded, expectedProject.SkillsNeeded)
// 	}

// 	for i, v := range result.SkillsNeeded {
// 		if v != expectedProject.SkillsNeeded[i] {
// 			t.Errorf("MapToQueryProject() SkillsNeeded = %v, want %v", result.SkillsNeeded, expectedProject.SkillsNeeded)
// 		}
// 	}
// }
