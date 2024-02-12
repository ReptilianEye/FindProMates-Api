package util_types

import "fmt"

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
func (s Skill) IsValid() error {
	switch s {
	case Python, Java, Go, JavaScript, React, Angular:
		return nil
	}
	return fmt.Errorf("%s is invalid skill", s)
}
