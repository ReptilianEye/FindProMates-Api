package collabrequests

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CollabRequestModel struct {
	C *mongo.Collection
}

const (
	ID          = "_id"
	ProjectID   = "project_id"
	RequesterID = "requester_id"
	ResponderID = "responder_id"
	Message     = "message"
	Feedback    = "feedback"
	Status      = "status"
)

var ctx = context.TODO()

func (m *CollabRequestModel) FindById(id primitive.ObjectID) (*CollabRequest, error) {
	var collabRequest CollabRequest
	err := m.C.FindOne(ctx, bson.M{ID: id}).Decode(&collabRequest)
	if err != nil {
		return nil, err
	}
	return &collabRequest, nil
}

func (m *CollabRequestModel) All() ([]*CollabRequest, error) {
	collabRequests := []*CollabRequest{}
	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &collabRequests)
	if err != nil {
		return nil, err
	}
	return collabRequests, nil
}

func (m *CollabRequestModel) AllByUser(userId primitive.ObjectID) ([]*CollabRequest, error) {
	collabRequests := []*CollabRequest{}
	cursor, err := m.C.Find(ctx, bson.M{RequesterID: userId})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &collabRequests)
	if err != nil {
		return nil, err
	}
	return collabRequests, nil
}
func (m *CollabRequestModel) AllByProject(projectId primitive.ObjectID) ([]*CollabRequest, error) {
	collabRequests := []*CollabRequest{}
	cursor, err := m.C.Find(ctx, bson.M{ProjectID: projectId})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &collabRequests)
	if err != nil {
		return nil, err
	}
	return collabRequests, nil
}
func (m *CollabRequestModel) Create(collabRequest *CollabRequest) error {
	result, err := m.C.InsertOne(ctx, collabRequest)
	if err != nil {
		return err
	}
	collabRequest.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (m *CollabRequestModel) Update(collabRequest *CollabRequest) error {
	_, err := m.C.ReplaceOne(ctx, bson.M{ID: collabRequest.ID}, collabRequest)
	return err
}

func (m *CollabRequestModel) Delete(id primitive.ObjectID) error {
	_, err := m.C.DeleteOne(ctx, bson.M{ID: id})
	return err
}
