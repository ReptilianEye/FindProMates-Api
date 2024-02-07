package users

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	C *mongo.Collection
}

// keys for params map
const (
	Id        string = "_id"
	FirstName string = "firstName"
	LastName  string = "lastName"
	Username  string = "username"
	Email     string = "email"
)

var ctx = context.TODO()

// All returns all users from the MongoDB collection.
func (m *UserModel) All() ([]User, error) {
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
func (m *UserModel) FindById(id primitive.ObjectID) (*User, error) {
	var user User
	err := m.C.FindOne((ctx), bson.M{Id: id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *UserModel) FindByUsername(username string) (primitive.ObjectID, error) {
	var user User
	err := m.C.FindOne(ctx, bson.M{Username: username}).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}
func (m *UserModel) FindByParameters(params map[string]string) ([]User, error) {
	query := bson.M{}
	for key, value := range params {
		if key == Id {
			panic("Use FindById to find by id")
		}
		query[key] = value
	}
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

func (m *UserModel) Create(user *User) (*User, error) {
	_, err := m.C.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserModel) Authenticate(username, password string) bool {
	var user User
	err := m.C.FindOne(ctx, bson.M{Username: username}).Decode(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	return checkPasswordHash(password, user.Password)

}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
