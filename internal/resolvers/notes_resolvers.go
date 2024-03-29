package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/notes"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNoteByStrId(id string) (*notes.Note, error) {
	noteId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.Notes.FindById(noteId)
}

func MapToQueryNote(note *notes.Note) *model.Note {
	return &model.Note{
		ID:           note.ID.Hex(),
		Project:      MapToQueryProject(ProjectByObjId(note.ProjectID)),
		AddedBy:      MapToQueryUser(UserByObjId(note.AddedBy)),
		Note:         note.Note,
		LastModified: note.LastModified,
	}
}
func MapToNoteFromNew(projectID primitive.ObjectID, addedBy primitive.ObjectID, note string) *notes.Note {
	return &notes.Note{
		ProjectID:    projectID,
		AddedBy:      addedBy,
		Note:         note,
		LastModified: time.Now(),
	}
}
func UpdateNote(note *notes.Note, newNote string) {
	note.Note = newNote
	note.LastModified = time.Now()
}
