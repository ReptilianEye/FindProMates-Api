package tasks

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskModel struct {
	C *mongo.Collection
}

const (
	Id               string = "_id"
	ProjectId        string = "project_id"
	TaskContent      string = "task"
	AddedBy          string = "added_by"
	AssignedTo       string = "assigned_to"
	CreatedAt        string = "created_at"
	Deadline         string = "deadline"
	PriorityLevel    string = "priority_level"
	CompletionStatus string = "completion_status"
)

var ctx = context.TODO()

func (m *TaskModel) AllAssignedToUser(userId primitive.ObjectID) ([]*Task, error) {
	tasks := []*Task{}
	cursor, err := m.C.Find(ctx, bson.D{
		{Key: AssignedTo, Value: userId},
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (m *TaskModel) FindById(id primitive.ObjectID) (*Task, error) {
	var task Task
	err := m.C.FindOne(ctx, bson.M{Id: id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *TaskModel) AllByProjectId(projectId primitive.ObjectID) ([]*Task, error) {
	tasks := []*Task{}
	cursor, err := m.C.Find(ctx, bson.M{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (m *TaskModel) Create(task *Task) error {
	result, err := m.C.InsertOne(ctx, task)
	if err != nil {
		return err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *TaskModel) Update(task *Task) error {
	_, err := m.C.ReplaceOne(ctx, bson.M{Id: task.ID}, task)
	return err
}

func (m *TaskModel) Delete(id primitive.ObjectID) error {
	_, err := m.C.DeleteOne(ctx, bson.M{Id: id})
	return err
}
