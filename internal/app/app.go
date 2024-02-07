package app

import (
	"example/FindProMates-Api/internal/database"
	"example/FindProMates-Api/internal/database/mongodb"
)

var App *Application

type Application struct {
	Projects *mongodb.ProjectModel
	Users    *mongodb.UserModel
}

func InitApp() {
	App = &Application{
		Projects: &mongodb.ProjectModel{
			C: database.Db.Collection("projects"),
		},
		Users: &mongodb.UserModel{
			C: database.Db.Collection("users"),
		},
	}
}
