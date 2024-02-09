package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/notes"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNoteById(id string) (*notes.Note, error) {
	noteId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.Notes.FindById(noteId)
}

func MapToQueryNote(note *notes.Note) *model.Note {
	return &model.Note{
		ID:        note.ID.Hex(),
		Project:   MapToQueryProject(ProjectByObjId(note.ProjectID)),
		AddedBy:   MapToQueryUser(UserByObjId(note.AddedBy)),
		Note:      note.Note,
		CreatedAt: note.CreatedAt,
	}
}
