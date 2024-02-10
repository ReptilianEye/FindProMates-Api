package util_types

import (
	"fmt"
	"log"
)

// types not stored in the database

type PriorityLevel struct {
	priority int
}

func (p PriorityLevel) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	}
	log.Fatal("Invalid Priority Level")
	return ""
}
func PriorityLevelFromString(s string) (PriorityLevel, error) {
	switch s {
	case "Low":
		return Low, nil
	case "Medium":
		return Medium, nil
	case "High":
		return High, nil
	}
	return PriorityLevel{}, fmt.Errorf("Invalid Priority Level")
}

var Low = PriorityLevel{priority: 1}
var Medium = PriorityLevel{priority: 2}
var High = PriorityLevel{priority: 3}
