package gsi

type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

type RoomHandler interface {
	OnInit(room Room)
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	Processor(room Room, action string, data map[string]interface{})
	OnJoined(room Room, user *User)
	OnLeft(room Room, user *User)
	OnReceived(room Room, user *User, message []byte)
	OnClose(room Room)
}
