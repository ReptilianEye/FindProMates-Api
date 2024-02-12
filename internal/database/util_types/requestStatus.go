package util_types

import "fmt"

type RequestStatus string

const (
	Pending  RequestStatus = "Pending"
	Accepted RequestStatus = "Accepted"
	Rejected RequestStatus = "Rejected"
)

func (r RequestStatus) String() string {
	return string(r)
}

func (r RequestStatus) IsValid() error {
	switch r {
	case Pending, Accepted, Rejected:
		return nil
	}
	return fmt.Errorf("%s is invalid request status", r)
}
