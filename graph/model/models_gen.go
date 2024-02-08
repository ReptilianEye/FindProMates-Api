// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Email    *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	Password string  `json:"password"`
}

type Mutation struct {
}

type NewProject struct {
	Name          string   `json:"name"`
	Description   *string  `json:"description,omitempty"`
	Collaborators []string `json:"collaborators,omitempty"`
}

type NewUser struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Username  *string  `json:"username,omitempty"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Skills    []string `json:"skills"`
}

type Project struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Owner         *User   `json:"owner"`
	Collaborators []*User `json:"collaborators"`
}

type Query struct {
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UpdateProject struct {
	Name          *string  `json:"name,omitempty"`
	Description   *string  `json:"description,omitempty"`
	Collaborators []string `json:"collaborators,omitempty"`
}

type User struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Skills    []string `json:"skills"`
}

type UserProjects struct {
	Owner        *User      `json:"owner"`
	Owned        []*Project `json:"owned"`
	Collaborated []*Project `json:"collaborated"`
}
