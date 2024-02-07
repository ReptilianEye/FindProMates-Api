package projects

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectModel struct {
	C *mongo.Collection
}

func (m *ProjectModel) All() ([]Project, error) {
	ctx := context.TODO()
	projects := []Project{}

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
func (m *ProjectModel) FindById(id string) (*Project, error) {
	ctx := context.TODO()
	project := Project{}
	err := m.C.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
func (m *ProjectModel) FindByOwner(owner string) ([]Project, error) {
	ctx := context.TODO()
	projects := []Project{}
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

func (m *ProjectModel) Insert(project Project) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), project)
}

func (m *ProjectModel) Update(id string, project Project) (*mongo.UpdateResult, error) {
	return m.C.ReplaceOne(context.TODO(), bson.M{"_id": id}, project)
}

func (m *ProjectModel) Delete(id string) (*mongo.DeleteResult, error) {
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": id})
}
