package app

import (
	database "example/FindProMates-Api/internal/db"
	"example/FindProMates-Api/internal/models/mongodb"
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
