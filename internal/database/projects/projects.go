package projects

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// keys for params map
const (
	ID               string = "_id"
	Name             string = "name"
	Owner            string = "owner"
	Description      string = "description"
	Public           string = "public"
	Collaborators    string = "collaborators"
	SkillsNeeded     string = "skills_needed"
	CompletionStatus string = "completion_status"
)

type ProjectModel struct {
	C *mongo.Collection
}

var ctx = context.TODO()

func (m *ProjectModel) All() ([]*Project, error) {
	projects := []*Project{}
	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (m *ProjectModel) FindById(id primitive.ObjectID) (*Project, error) {
	var project Project
	err := m.C.FindOne(ctx, bson.M{ID: id}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
func (m *ProjectModel) FindByOwner(owner primitive.ObjectID) ([]*Project, error) {
	projects := []*Project{}
	cursor, err := m.C.Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (m *ProjectModel) FindAllUserIsCollaborator(user primitive.ObjectID) ([]*Project, error) {
	var projects []*Project
	cursor, err := m.C.Find(ctx, bson.D{
		{Key: Collaborators, Value: user}, //https://www.mongodb.com/docs/manual/tutorial/query-arrays/
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil

}

func (m *ProjectModel) Create(project *Project) error {
	result, err := m.C.InsertOne(context.TODO(), project)
	if err != nil {
		return err
	}
	project.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *ProjectModel) Update(id primitive.ObjectID, project *Project) error {
	_, err := m.C.ReplaceOne(context.TODO(), bson.M{ID: id}, project)
	return err
}

func (m *ProjectModel) Delete(id primitive.ObjectID) error {
	_, err := m.C.DeleteOne(context.TODO(), bson.M{ID: id})
	return err
}
