package collabrequest

import (
	"example/FindProMates-Api/internal/database/util_types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollabRequestCollection string = "collab_requests"

type CollabRequest struct {
	ID          primitive.ObjectID       `bson:"_id,omitempty"`
	ProjectID   primitive.ObjectID       `bson:"project_id,omitempty"`
	RequesterID primitive.ObjectID       `bson:"requester_id,omitempty"`
	ResponderID primitive.ObjectID       `bson:"responder_id,omitempty"`
	Message     string                   `bson:"message"`
	Feedback    string                   `bson:"feedback"`
	Status      util_types.RequestStatus `bson:"status,omitempty"`
}
