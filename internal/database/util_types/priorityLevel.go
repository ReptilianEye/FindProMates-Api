package util_types

import "fmt"

type PriorityLevel string

const (
	LowPriority    PriorityLevel = "Low"
	MediumPriority PriorityLevel = "Medium"
	HighPriority   PriorityLevel = "High"
)

func (p PriorityLevel) String() string {
	return string(p)
}

func (c PriorityLevel) IsValid() error {
	switch c {
	case LowPriority, MediumPriority, HighPriority:
		return nil
	}
	return fmt.Errorf("%s is invalid priority level", c)
}
