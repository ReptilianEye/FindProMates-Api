package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"example/FindProMates-Api/graph/model"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/auth"
	"example/FindProMates-Api/internal/pkg/jwt"
	"example/FindProMates-Api/internal/pkg/utils"
	"example/FindProMates-Api/internal/resolvers"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := resolvers.MapToUser(input)
	_, err := app.App.Users.Create(&user)
	if err != nil {
		return nil, err
	}
	return resolvers.MapToQueryUser(user), nil
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	userId := auth.ForContext(ctx)
	if userId == "" {
		return nil, fmt.Errorf("access denied")
	}
	ownerId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	project := resolvers.MapToProject(input, ownerId)
	_, err = app.App.Projects.Create(&project)
	if err != nil {
		return nil, err
	}
	return resolvers.MapToQueryProject(project), nil
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: UpdateProject - updateProject"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	username := input.Username
	password := input.Password
	if !app.App.Users.Authenticate(username, password) {
		return "", fmt.Errorf("username or password is incorrect")
	}
	id, err := app.App.Users.FindByUsername(username)
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(id.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	usersArr, err := app.App.Users.All()
	if err != nil {
		return nil, err
	}
	return utils.MapTo(usersArr, resolvers.MapToQueryUser), nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	projectsArr, err := app.App.Projects.All()
	if err != nil {
		return nil, err
	}
	return utils.MapTo(projectsArr, resolvers.MapToQueryProject), nil
}

// ProjectsByUser is the resolver for the projectsByUser field.
func (r *queryResolver) ProjectsByUser(ctx context.Context, id string) (*model.UserProjects, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	owner, err := app.App.Users.FindById(idObj)
	if err != nil {
		return nil, err
	}

	projects, err := app.App.Projects.FindByOwner(idObj)
	if err != nil {
		return nil, err
	}

	collaboratedProjects, err := app.App.Projects.FindAllUserIsCollaborator(idObj)
	if err != nil {
		return nil, err
	}
	return &model.UserProjects{
		Owner:        resolvers.MapToQueryUser(*owner),
		Owned:        utils.MapTo(projects, resolvers.MapToQueryProject),
		Collaborated: utils.MapTo(collaboratedProjects, resolvers.MapToQueryProject),
	}, nil
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	project, err := app.App.Projects.FindById(idObj)
	if err != nil {
		return nil, err
	}
	return resolvers.MapToQueryProject(*project), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
