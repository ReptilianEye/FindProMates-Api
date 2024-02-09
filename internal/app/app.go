package app

import (
	"context"
	"example/FindProMates-Api/internal/database"
	"example/FindProMates-Api/internal/database/notes"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/tasks"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/pkg/jwt"

	"github.com/joho/godotenv"
)

var App *Application

type Application struct {
	Projects *projects.ProjectModel
	Users    *users.UserModel
	Notes    *notes.NoteModel
	Tasks    *tasks.TaskModel
}

func InitApp() context.CancelFunc {
	jwtDone := make(chan bool)
	dbCancel := make(chan context.CancelFunc)
	godotenv.Load()
	go func() {
		jwt.InitJWT()
		jwtDone <- true
	}()
	go func() {
		cancel := database.InitDB()
		App = &Application{
			Projects: &projects.ProjectModel{
				C: database.Db.Collection("projects"),
			},
			Users: &users.UserModel{
				C: database.Db.Collection("users"),
			},
			Notes: &notes.NoteModel{
				C: database.Db.Collection("notes"),
			},
			Tasks: &tasks.TaskModel{
				C: database.Db.Collection("tasks"),
			},
		}
		dbCancel <- cancel
	}()
	<-jwtDone
	return <-dbCancel
}
