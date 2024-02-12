package util_types

import "fmt"

type CompletionStatus string

const (
	NotStarted CompletionStatus = "Not Started"
	InProgress CompletionStatus = "In Progress"
	Completed  CompletionStatus = "Completed"
)

func (c CompletionStatus) String() string {
	return string(c)
}

func (c CompletionStatus) IsValid() error {
	switch c {
	case NotStarted, InProgress, Completed:
		return nil
	}
	return fmt.Errorf("%s is invalid completion status", c)
}
