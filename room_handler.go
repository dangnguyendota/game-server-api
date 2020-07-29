package gsi

type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

type RoomData struct {
	User *User
	Data []byte
	Time int64
}

type RoomHandler interface {
	OnInit(room Room)
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	Processor(room Room, action string, data map[string]interface{})
	OnJoined(room Room, user *User)
	OnLeft(room Room, user *User)
	Loop(room Room, roomChan <-chan *RoomData)
	OnClose(room Room)
}
