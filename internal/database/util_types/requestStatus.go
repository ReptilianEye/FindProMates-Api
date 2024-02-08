package util_types

type RequestStatus string

const (
	Pending  RequestStatus = "Pending"
	Accepted RequestStatus = "Accepted"
	Rejected RequestStatus = "Rejected"
)

func (c RequestStatus) String() string {
	return string(c)
}

func (c RequestStatus) IsValid() bool {
	switch c {
	case Pending, Accepted, Rejected:
		return true
	}
	return false
}
