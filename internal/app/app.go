package app

import (
	"context"
	"example/FindProMates-Api/internal/database"
	"example/FindProMates-Api/internal/database/collabrequest"
	"example/FindProMates-Api/internal/database/notes"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/tasks"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/pkg/jwt"

	"github.com/joho/godotenv"
)

var App *Application

type Application struct {
	Projects       *projects.ProjectModel
	Users          *users.UserModel
	Notes          *notes.NoteModel
	Tasks          *tasks.TaskModel
	CollabRequests *collabrequest.CollabRequestModel
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
				C: database.Db.Collection(projects.ProjectCollection),
			},
			Users: &users.UserModel{
				C: database.Db.Collection(users.UserCollection),
			},
			Notes: &notes.NoteModel{
				C: database.Db.Collection(notes.NoteCollection),
			},
			Tasks: &tasks.TaskModel{
				C: database.Db.Collection(tasks.TaskCollection),
			},
			CollabRequests: &collabrequest.CollabRequestModel{
				C: database.Db.Collection(collabrequest.CollabRequestCollection),
			},
		}
		dbCancel <- cancel
	}()
	<-jwtDone
	return <-dbCancel
}
