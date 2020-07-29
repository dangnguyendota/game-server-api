package gsi

type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

type RoomHandler interface {
	OnInit(room Room)
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	Processor(room Room, action string, data map[string]interface{})
	onJoined(room Room, user *User)
	onLeft(room Room, user *User)
	onMessage(room Room, user *User, message []byte)
	OnClose(room Room)
}
