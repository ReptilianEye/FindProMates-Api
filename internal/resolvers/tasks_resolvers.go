package resolvers

import (
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/database/projects"
	"example/FindProMates-Api/internal/database/tasks"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"fmt"
	"slices"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TaskByStrId(id string) (*tasks.Task, error) {
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
	fmt.Println(task)
	assignedTo := utils.MapTo(task.AssignedTo, UserByObjId)
	return &model.Task{
		ID:               task.ID.Hex(),
		Project:          MapToQueryProject(ProjectByObjId(task.ProjectID)),
		AddedBy:          MapToQueryUser(UserByObjId(task.AddedBy)),
		AssignedTo:       utils.MapTo(assignedTo, MapToQueryUser),
		Task:             task.Task,
		LastModified:     task.LastModified,
		Deadline:         &task.Deadline,
		PriorityLevel:    task.PriorityLevel.String(),
		CompletionStatus: task.CompletionStatus.String(),
	}
}
func MapToTaskFromNew(projectID primitive.ObjectID, addedBy primitive.ObjectID, task model.NewTask) (*tasks.Task, error) {
	priority := util_types.MediumPriority
	if task.PriorityLevel != nil {
		priority = util_types.PriorityLevel(*task.PriorityLevel)
		if err := priority.IsValid(); err != nil {
			return nil, err
		}
	}
	var assignedTo []primitive.ObjectID
	if task.AssignedTo != nil {
		handled, err := handleAssignedTo(task.AssignedTo)
		if err != nil {
			return nil, err
		}
		assignedTo = handled
	}
	return &tasks.Task{
		ProjectID:        projectID,
		AddedBy:          addedBy,
		AssignedTo:       assignedTo,
		Task:             task.Task,
		Deadline:         utils.Elivis(task.Deadline, time.Now().AddDate(10, 0, 0)),
		LastModified:     time.Now(),
		PriorityLevel:    priority,
		CompletionStatus: utils.Ternary(len(assignedTo) > 0, util_types.InProgress, util_types.NotStarted).(util_types.CompletionStatus),
	}, nil
}
func UpdateTask(task *tasks.Task, updatedTask model.UpdatedTask) error {
	task.Task = utils.Elivis(updatedTask.Task, task.Task)
	task.Deadline = utils.Elivis(updatedTask.Deadline, task.Deadline)
	if updatedTask.PriorityLevel != nil {
		task.PriorityLevel = util_types.PriorityLevel(*updatedTask.PriorityLevel)
		if err := task.PriorityLevel.IsValid(); err != nil {
			return err
		}
	}
	if updatedTask.CompletionStatus != nil {
		task.CompletionStatus = util_types.CompletionStatus(*updatedTask.CompletionStatus)
		if err := task.CompletionStatus.IsValid(); err != nil {
			return err
		}
	}
	if updatedTask.AssignedTo != nil {
		handled, err := handleAssignedTo(updatedTask.AssignedTo, task.AssignedTo...)
		if err != nil {
			return err
		}
		project := ProjectByObjId(task.ProjectID)
		for _, id := range handled {
			if !slices.Contains(project.Collaborators, id) {
				return fmt.Errorf("user %s is not a collaborator", id.Hex())
			}
		}
		if len(handled) > 0 && task.CompletionStatus == util_types.NotStarted {
			task.CompletionStatus = util_types.InProgress
		}
		task.AssignedTo = handled
	}
	return nil
}

func handleAssignedTo(assignedTo []string, base ...primitive.ObjectID) ([]primitive.ObjectID, error) {
	return utils.MergeSlices(base, assignedTo, utils.SafeIDToString, utils.SafeStringToID)
}
