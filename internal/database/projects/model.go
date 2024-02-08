package projects

import (
	"example/FindProMates-Api/internal/database/util_types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ProjectCollection string = "projects"

type Project struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Owner         primitive.ObjectID   `bson:"owner,omitempty"`
	Name          string               `bson:"name,omitempty"`
	Description   string               `bson:"description,omitempty"`
	Collaborators []primitive.ObjectID `bson:"collaborators,omitempty"`
	SkillsNeeded  []util_types.Skill   `bson:"skills_needed,omitempty"`
}
