package dbschema

const UserCollection string = "users"

type Skill string

const (
	Python     Skill = "Python"
	Java       Skill = "Java"
	Go         Skill = "Go"
	JavaScript Skill = "JavaScript"
	React      Skill = "React"
	Angular    Skill = "Angular"
)

type User struct {
	FirstName string  `bson:"first_name,omitempty" `
	LastName  string  `bson:"last_name,omitempty"`
	Username  string  `bson:"username,omitempty"`
	Email     string  `bson:"email,omitempty"`
	Password  string  `bson:"password,omitempty"`
	Skills    []Skill `bson:"interests"`
}
