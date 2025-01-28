package entity

type UserType int

const (
	ADMIN UserType = iota
	AGENT
	CUSTOMER
)

func (u UserType) String() string {
	switch u {
	case ADMIN:
		return "ADMIN"
	case AGENT:
		return "AGENT"
	case CUSTOMER:
		return "CUSTOMER"
	default:
		return "UNKNOWN"
	}
}

type User struct {
	Id         int      `json:"id"`
	ExternalId string   `json:"externalId"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Type       UserType `json:"type"`
}
