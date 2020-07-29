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
	// init the game and return game state
	OnInit(room Room) interface{}
	// check join condition for allow or disallow a user join this room
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	// when scheduled an action, this action will jump here
	Processor(room Room, action string, data map[string]interface{})
	// when a user joined this room (players only) and return the game state
	OnJoined(room Room, user *User) interface{}
	// when a player left room (or disconnected) (players only) and return the game state
	OnLeft(room Room, user *User) interface{}
	// handle all queued messages from clients and return the game state
	Loop(room Room, roomChan <-chan *RoomData) interface{}
	// when room is closed
	OnClose(room Room)
}
