package main

import (
	"example/FindProMates-Api/internal/models/mongodb"

	"github.com/graphql-go/graphql"
)

func (app *Application) GetUsersSchema() *graphql.Schema {

	var usersQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(usersType),
				Description: "Get all users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
					// return app.Users.All()
				},
			},
			"user": &graphql.Field{
				Type:        usersType,
				Description: "Get user by username or email or first name or last name",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var query = map[string]string{}

					if username := p.Args["username"].(string); username != "" {
						query[mongodb.Username] = username
					}
					if email := p.Args["email"].(string); email != "" {
						query[mongodb.Email] = email
					}
					if firstName := p.Args["firstName"].(string); firstName != "" {
						query[mongodb.FirstName] = firstName
					}
					if lastName := p.Args["lastName"].(string); lastName != "" {
						query[mongodb.LastName] = lastName
					}
					return app.Users.FindByParameters(query)
				},
			},
		},
	})

	UsersSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: usersQuery,
	})
	return &UsersSchema
}

var usersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"skills": &graphql.Field{
			Type: graphql.NewList(skillType),
		},
	},
})
var skillType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Skill",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// var UsersSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
// 	Query: usersQuery,
// })

// var usersQuery = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "RootQuery",
// 	Fields: graphql.Fields{
// 		"users": &graphql.Field{
// 			Type:        graphql.NewList(usersType),
// 			Description: "Get all users",
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return app.Users.All()
// 			},
// 		},
// 		"user": &graphql.Field{
// 			Type:        usersType,
// 			Description: "Get user by username or email or first name or last name",
// 			Args: graphql.FieldConfigArgument{
// 				"username": &graphql.ArgumentConfig{
// 					Type: graphql.String,
// 				},
// 				"email": &graphql.ArgumentConfig{
// 					Type: graphql.String,
// 				},
// 				"firstName": &graphql.ArgumentConfig{
// 					Type: graphql.String,
// 				},
// 				"lastName": &graphql.ArgumentConfig{
// 					Type: graphql.String,
// 				},
// 			},
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				var query = map[string]string{}

// 				if username := p.Args["username"].(string); username != "" {
// 					query[mongodb.Username] = username
// 				}
// 				if email := p.Args["email"].(string); email != "" {
// 					query[mongodb.Email] = email
// 				}
// 				if firstName := p.Args["firstName"].(string); firstName != "" {
// 					query[mongodb.FirstName] = firstName
// 				}
// 				if lastName := p.Args["lastName"].(string); lastName != "" {
// 					query[mongodb.LastName] = lastName
// 				}
// 				return app.Users.FindByParameters(query)
// 			},
// 		},
// 	},
// })

// var usersType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "User",
// 	Fields: graphql.Fields{
// 		"firstName": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"lastName": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"username": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"email": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"password": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"skills": &graphql.Field{
// 			Type: graphql.NewList(skillType),
// 		},
// 	},
// })
// var skillType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Skill",
// 	Fields: graphql.Fields{
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 	},
// })
