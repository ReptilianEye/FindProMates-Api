package users

import (
	"context"
	"example/FindProMates-Api/internal/pkg/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo map[string]string

type UserModel struct {
	C *mongo.Collection
}

// keys for params map
const (
	ID        string = "_id"
	FirstName string = "firstName"
	LastName  string = "lastName"
	Username  string = "username"
	Email     string = "email"
	Skills    string = "skills"
)

var ctx = context.TODO()

// All returns all users from the MongoDB collection.
func (m *UserModel) All() ([]*User, error) {
	users := []*User{}

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
	err := m.C.FindOne((ctx), bson.M{ID: id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *UserModel) FindByUserInfo(usernameOrEmail UserInfo) (*User, error) {
	var user User
	err := m.C.FindOne(ctx, usernameOrEmail).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *UserModel) FindByParameters(params map[string]string) ([]User, error) {
	query := bson.M{}
	for key, value := range params {
		if key == ID {
			log.Panic("use FindById to find by id")
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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	if err := validateUserInfo(user); err != nil {
		return nil, err
	}
	result, err := m.C.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *UserModel) Update(user *User, changingPassword bool) (*mongo.UpdateResult, error) {
	if changingPassword {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}
	if err := validateUserInfo(user); err != nil {
		return nil, err
	}

	return m.C.ReplaceOne(ctx, bson.M{ID: user.ID}, user)
}

func (m *UserModel) Authenticate(userInfo UserInfo, password string) bool {
	user, err := m.FindByUserInfo(userInfo)
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

func BuildUserInfo(username, email *string) UserInfo {
	var usersInfo = make(map[string]string)
	if username != nil {
		usersInfo[Username] = *username
	}
	if email != nil {
		usersInfo[Email] = *email
	}
	return usersInfo
}
func validateUserInfo(user *User) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		return err
	}
	if name, err := utils.ValidateName(user.FirstName); err != nil {
		return err
	} else {
		user.FirstName = name
	}
	if name, err := utils.ValidateName(user.LastName); err != nil {
		return err
	} else {
		user.LastName = name
	}
	return nil
}
