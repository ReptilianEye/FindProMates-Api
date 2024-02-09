package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/tasks"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TaskById(id string) (*tasks.Task, error) {
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.Tasks.FindById(taskId)
}
func TasksAssignedToUser(user *users.User) ([]*model.Task, error) {
	tasks, err := app.App.Tasks.AllAssignedToUser(user.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(tasks, MapToQueryTask), nil
}
func TasksByProject(project *projects.Project) ([]*model.Task, error) {
	tasks, err := app.App.Tasks.AllByProjectId(project.ID)
	if err != nil {
		return nil, err
	}
	return utils.MapTo(tasks, MapToQueryTask), nil
}

func MapToQueryTask(task *tasks.Task) *model.Task {
	assignedTo := utils.MapTo(task.AssignedTo, UserByObjId)
	return &model.Task{
		ID:               task.ID.Hex(),
		Project:          MapToQueryProject(ProjectByObjId(task.ProjectID)),
		AddedBy:          MapToQueryUser(UserByObjId(task.AddedBy)),
		AssignedTo:       utils.MapTo(assignedTo, MapToQueryUser),
		Task:             task.Task,
		CreatedAt:        task.CreatedAt,
		Deadline:         &task.Deadline,
		PriorityLevel:    task.PriorityLevel.String(),
		CompletionStatus: task.CompletionStatus.String(),
	}
}
