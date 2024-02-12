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
	var responder *model.User
	if !collabRequest.ResponderID.IsZero() {
		responder = MapToQueryUser(UserByObjId(collabRequest.ResponderID))
	}
	var feedback *string
	if collabRequest.Feedback != "" {
		feedback = &collabRequest.Feedback
	}
	return &model.CollabRequest{
		ID:        collabRequest.ID.Hex(),
		Project:   MapToQueryProject(ProjectByObjId(collabRequest.ProjectID)),
		Requester: MapToQueryUser(UserByObjId(collabRequest.RequesterID)),
		Responder: responder,
		Message:   collabRequest.Message,
		Feedback:  feedback,
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
func UpdateCollabRequest(collabReq *collabrequests.CollabRequest, responder *users.User, status string, feedback string) error {
	if collabReq.Status == util_types.Accepted || collabReq.Status == util_types.Rejected {
		return fmt.Errorf("cannot update request that has been accepted or rejected")
	}
	collabReq.Status = util_types.RequestStatus(status)
	if err := collabReq.Status.IsValid(); err != nil {
		return err
	}
	if collabReq.Status == util_types.Pending {
		return fmt.Errorf("cannot update request to pending")
	}
	if len(feedback) < 1 && collabReq.Status == util_types.Rejected {
		return fmt.Errorf("feedback cannot be empty")
	}
	collabReq.Feedback = feedback
	collabReq.ResponderID = responder.ID
	return nil
}
