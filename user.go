package api

import (
	"github.com/google/uuid"
)

type User struct {
	Sid         uuid.UUID         `json:"sid"` // session id
	Id          uuid.UUID         `json:"id"`  // user id
	Username    string            `json:"username"`
	DisplayName string            `json:"display_name"`
	Avatar      string            `json:"avatar"`
	Attributes  map[string]string `json:"attributes"`
}
