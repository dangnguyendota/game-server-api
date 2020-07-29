package gsi

import (
	"github.com/google/uuid"
)

type User struct {
	// base
	SID         uuid.UUID `json:"sid"` // session id
	ID          uuid.UUID `json:"id"`  // user id
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar"`
	// add your new attributes here
	Attributes map[string]string `json:"attributes"`
}

func (user *User) UpdateAttribute(attribute string, val string) {
	if _, ok := user.Attributes[attribute]; !ok {
		return
	}
	user.Attributes[attribute] = val
}
