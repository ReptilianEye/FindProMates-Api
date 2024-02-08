package util_types

import "log"

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

var Low = PriorityLevel{priority: 1}
var Medium = PriorityLevel{priority: 2}
var High = PriorityLevel{priority: 3}
