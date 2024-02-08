package util_types

type CompletionStatus string

const (
	NotStarted CompletionStatus = "Not Started"
	InProgress CompletionStatus = "In Progress"
	Completed  CompletionStatus = "Completed"
)

func (c CompletionStatus) String() string {
	return string(c)
}

func (c CompletionStatus) IsValid() bool {
	switch c {
	case NotStarted, InProgress, Completed:
		return true
	}
	return false
}
