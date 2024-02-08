package tasks

import (
	"example/FindProMates-Api/internal/database/util_types"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TaskCollection string = "tasks"

type Task struct {
	ID               primitive.ObjectID          `bson:"_id,omitempty"`
	ProjectID        primitive.ObjectID          `bson:"project_id,omitempty"`
	Name             string                      `bson:"name,omitempty"`
	Task             string                      `bson:"task,omitempty"`
	AddedBy          primitive.ObjectID          `bson:"added_by,omitempty"`
	CreatedAt        time.Time                   `bson:"created_at,omitempty"`
	Deadline         time.Time                   `bson:"deadline,omitempty"`
	PriorityLevel    util_types.PriorityLevel    `bson:"priority_level,omitempty"`
	CompletionStatus util_types.CompletionStatus `bson:"completion_status,omitempty"`
}
