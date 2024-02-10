package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/collabrequests"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CollabRequestByStrId(id string) (*collabrequests.CollabRequest, error) {
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

func MapToQueryCollabRequest(collabRequest *collabrequests.CollabRequest) *model.CollabRequest {
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
func MapToCollabRequestFromNew(projectId primitive.ObjectID, requesterId primitive.ObjectID, message string) *collabrequests.CollabRequest {
	return &collabrequests.CollabRequest{
		ProjectID:   projectId,
		RequesterID: requesterId,
		Message:     message,
		Status:      util_types.Pending,
	}
}
func UpdateCollabRequest(collabReq *collabrequests.CollabRequest, status string, feedback string) error {
	statusT := util_types.RequestStatus(status)
	if !statusT.IsValid() {
		return fmt.Errorf(`Invalid status: "%s"`, status)
	}
	collabReq.Status = statusT
	collabReq.Feedback = feedback
	return nil
}
