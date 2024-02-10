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
	LastMod     string = "last_modified"
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

func (m *NoteModel) Create(note *Note) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(ctx, note)
}

func (m *NoteModel) Update(note *Note) (*mongo.UpdateResult, error) {
	return m.C.UpdateOne(ctx, bson.M{Id: note.ID}, bson.M{"$set": note})
}
func (m *NoteModel) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return m.C.DeleteOne(ctx, bson.M{Id: id})
}
