package notes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteModel struct {
	C *mongo.Collection
}

const (
	Id          string = "_id"
	ProjectId   string = "project_id"
	AddedBy     string = "added_by"
	NoteContent string = "note"
	CreatedAt   string = "created_at"
)

var ctx = context.TODO()

func (m *NoteModel) FindById(id primitive.ObjectID) (*Note, error) {
	var note Note
	err := m.C.FindOne(ctx, bson.M{Id: id}).Decode(&note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (m *NoteModel) AllByProjectId(projectId primitive.ObjectID) ([]*Note, error) {
	notes := []*Note{}
	cursor, err := m.C.Find(ctx, bson.M{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
