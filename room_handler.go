package gs_interface

import "github.com/google/uuid"

type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

type RoomDataMessage struct {
	From uuid.UUID
	To   []uuid.UUID
	Code int64
	Data []byte
	Time int64
}

type RoomHandler interface {
	OnInit(room Room)
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	Processor(room Room, action string, data map[string]interface{})
	HandleJoin(room Room, user *User)
	HandleLeave(room Room, user *User)
	HandleData(room Room, message *RoomDataMessage)
	OnClose(room Room)
}