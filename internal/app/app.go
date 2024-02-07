package app

import (
	"example/FindProMates-Api/internal/database"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/users"
)

var App *Application

type Application struct {
	Projects *projects.ProjectModel
	Users    *users.UserModel
}

func InitApp() {
	App = &Application{
		Projects: &projects.ProjectModel{
			C: database.Db.Collection("projects"),
		},
		Users: &users.UserModel{
			C: database.Db.Collection("users"),
		},
	}
}
