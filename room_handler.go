package api

type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

type RoomData struct {
	User *User
	Data []byte
	Time int64
}

// must implement for handle room messages
type RoomHandler interface {
	// reload phòng từ data lưu trong redis sau khi khởi động lại server
	OnReload(room Room, state string) interface{}
	// khởi tạo phòng
	OnInit(room Room) interface{}
	// kiểm tra điều kiện tham gia phòng của người chơi
	AllowJoin(room Room, user *User) *CheckJoinConditionResult
	// sau khi một action đã được schedule trước đó
	Processor(room Room, action string, data map[string]interface{}) interface{}
	// when a user joined this room (players only) and return the game state
	OnJoined(room Room, user *User) interface{}
	// when a player left room (or disconnected) (players only) and return the game state
	OnLeft(room Room, user *User) interface{}
	// handle all queued messages from clients and return the game state
	Loop(room Room, roomChan <-chan *RoomData) interface{}
	// when room is closed
	OnClose(room Room)
}

// create a new room handler
type CreateRoomHandlerFunc func() RoomHandler
