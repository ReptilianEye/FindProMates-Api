package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/collabrequest"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CollabRequestByStrId(id string) (*collabrequest.CollabRequest, error) {
	collabReqId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.CollabRequests.FindById(collabReqId)
}

func CollabRequestsByUser(user *users.User) ([]*model.CollabRequest, error) {
	collabRequests, err := app.App.CollabRequests.AllByUser(user.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(collabRequests, MapToQueryCollabRequest), nil
}

func CollabRequestsByProject(project *projects.Project) ([]*model.CollabRequest, error) {
	collabRequests, err := app.App.CollabRequests.AllByProject(project.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(collabRequests, MapToQueryCollabRequest), nil
}

func MapToQueryCollabRequest(collabRequest *collabrequest.CollabRequest) *model.CollabRequest {
	return &model.CollabRequest{
		ID:        collabRequest.ID.Hex(),
		Project:   MapToQueryProject(ProjectByObjId(collabRequest.ProjectID)),
		Requester: MapToQueryUser(UserByObjId(collabRequest.RequesterID)),
		Responder: MapToQueryUser(UserByObjId(collabRequest.ResponderID)),
		Message:   collabRequest.Message,
		Feedback:  &collabRequest.Feedback,
		Status:    collabRequest.Status.String(),
	}
}
