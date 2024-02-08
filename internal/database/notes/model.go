package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const NoteCollection string = "notes"

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `bson:"project_id,omitempty"`
	AddedBy   primitive.ObjectID `bson:"added_by,omitempty"`
	Note      string             `bson:"note,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
