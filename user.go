package gs_interface

import (
	"./proto"
	"github.com/google/uuid"
)

type User struct {
	// base
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar"`
	// add your new attributes here
	Attributes map[string]string `json:"attributes"`
}

func (user *User) ConvertToApi() *ip.Player {
	player := &ip.Player{
		Id:          user.ID.String(),
		Name:        user.Username,
		DisplayName: user.DisplayName,
		Avatar:      user.Avatar,
		Attributes:  user.Attributes,
	}
	if player.Attributes != nil {
		if sid, ok := player.Attributes["sid"]; ok {
			player.Sid = sid
		}
	}
	return player
}

func (user *User) UpdateAttribute(attribute string, val string) {
	if _, ok := user.Attributes[attribute]; !ok {
		return
	}
	user.Attributes[attribute] = val
}
