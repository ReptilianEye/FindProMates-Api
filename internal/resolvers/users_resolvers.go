package resolvers

import (
	"context"
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/auth"
	"example/FindProMates-Api/internal/database/users"
	"example/FindProMates-Api/internal/database/util_types"
	"example/FindProMates-Api/internal/pkg/utils"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserByStrId(id string) (*users.User, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return app.App.Users.FindById(userId)
}

// when we are sure that the user exists
func UserByObjId(id primitive.ObjectID) *users.User {
	user, err := app.App.Users.FindById(id)
	if err != nil {
		log.Panic(err)
	}
	return user
}
func MapToQueryUser(user *users.User) *model.User {
	return &model.User{
		ID:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Skills: utils.MapTo(user.Skills, func(skill util_types.Skill) string {
			return skill.String()
		}),
	}
}
func MapToUser(user model.NewUser) *users.User {
	return &users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  utils.Elivis(user.Username, utils.CreateUsername(user.FirstName, user.LastName)),
		Email:     user.Email,
		Password:  user.Password,
		Skills: utils.MapTo(user.Skills, func(skill string) util_types.Skill {
			return util_types.Skill(skill)
		}),
		// Projects: make([]projects.Project, 0),
	}
}
func Authenticate(username, email *string, password string) (*users.User, error) {
	userAuthInfo := users.BuildUserInfo(username, email)
	if len(userAuthInfo) == 0 {
		return nil, fmt.Errorf("username or email is required")
	}
	if len(userAuthInfo) == 2 {
		return nil, fmt.Errorf("username and email are not allowed at the same time")
	}
	if !app.App.Users.Authenticate(userAuthInfo, password) {
		return nil, fmt.Errorf("username or password is incorrect")
	}
	user, err := app.App.Users.FindByUserInfo(userAuthInfo)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func UserFromContex(ctx context.Context) (*users.User, error) {
	userId := auth.ForContext(ctx)
	if userId == "" {
		return nil, fmt.Errorf("access denied")
	}
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	user, err := app.App.Users.FindById(userIdObj)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func UpdateUser(base *users.User, user model.UpdatedUser) {
	base.FirstName = utils.Elivis(user.FirstName, base.FirstName)
	base.LastName = utils.Elivis(user.LastName, base.LastName)
	base.Username = utils.Elivis(user.Username, base.Username)
	base.Email = utils.Elivis(user.Email, base.Email)
	base.Password = utils.Elivis(user.NewPassword, base.Password)
	base.Skills = utils.MapTo(user.Skills, func(skill string) util_types.Skill {
		return util_types.Skill(skill)
	})
}
