package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	C *mongo.Collection
}

// keys for params map
const (
	Id        string = "_id"
	FirstName string = "first_name"
	LastName  string = "last_name"
	Username  string = "username"
	Email     string = "email"
	Password  string = "password"
)

// All returns all users from the MongoDB collection.
func (m *UserModel) All() ([]User, error) {
	ctx := context.TODO()
	users := []User{}

	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (m *UserModel) FindById(id string) (*User, error) {
	users, err := m.FindByParameters(map[string]string{Id: id})
	if err != nil {
		return nil, err
	}
	return &users[0], nil
}

func (m *UserModel) FindByParameters(params map[string]string) ([]User, error) {
	query := bson.M{}
	for key, value := range params {
		query[key] = value
	}
	ctx := context.TODO()
	cursor, err := m.C.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	users := []User{}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
