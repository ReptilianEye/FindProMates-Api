package app

import (
	"context"
	"example/FindProMates-Api/internal/database"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/pkg/jwt"

	"github.com/joho/godotenv"
)

var App *Application

type Application struct {
	Projects *projects.ProjectModel
	Users    *users.UserModel
}

func InitApp() context.CancelFunc {
	jwtDone := make(chan bool)
	godotenv.Load()
	go func() {
		jwt.InitJWT()
		jwtDone <- true
	}()
	cancel := database.InitDB()
	App = &Application{
		Projects: &projects.ProjectModel{
			C: database.Db.Collection("projects"),
		},
		Users: &users.UserModel{
			C: database.Db.Collection("users"),
		},
	}
	<-jwtDone
	return cancel
}
