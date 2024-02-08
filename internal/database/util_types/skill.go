package util_types

//types not stored in the database
type Skill string

const (
	Python     Skill = "Python"
	Java       Skill = "Java"
	Go         Skill = "Go"
	JavaScript Skill = "JavaScript"
	React      Skill = "React"
	Angular    Skill = "Angular"
)

func (s Skill) String() string {
	return string(s)
}
func (s Skill) IsValid() bool {
	switch s {
	case Python, Java, Go, JavaScript, React, Angular:
		return true
	}
	return false
}
