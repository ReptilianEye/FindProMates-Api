package projects

import (
	"example/FindProMates-Api/internal/database/util_types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ProjectCollection string = "projects"

type Project struct {
	ID               primitive.ObjectID          `bson:"_id,omitempty"`
	Owner            primitive.ObjectID          `bson:"owner,omitempty"`
	Name             string                      `bson:"name,omitempty"`
	Description      string                      `bson:"description,omitempty"`
	Public           bool                        `bson:"public"`
	CompletionStatus util_types.CompletionStatus `bson:"completion_status,omitempty"`
	Collaborators    []primitive.ObjectID        `bson:"collaborators,omitempty"`
	SkillsNeeded     []util_types.Skill          `bson:"skills_needed,omitempty"`

	// RequestsToCollab []primitive.ObjectID `bson:"requests_to_collab,omitempty"` //joined from ColabRequests
}
