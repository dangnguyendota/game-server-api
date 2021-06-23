package api

import (
	"github.com/google/uuid"
	"sync"
)

var (
	handlersMu sync.RWMutex
	handlers = make(map[string]CreateRoomHandlerFunc)
)

// điều kiện tham gia phòng của người chơi trả về
// Allow bằng true nếu cho phép tham gia phòng
// nếu Allow bằng false thì Reason phải khác rỗng mô tả lý do
// không cho người chơi tham gia phòng
type CheckJoinConditionResult struct {
	Allow  bool
	Reason string
}

// dữ liệu người chơi gửi lên server sẽ được convert sang dạng này
// UserId là id của người chơi.
// Data là dữ liệu người chơi gửi lên có thể dưới dạng proto hoặc json
// hoặc bất cứ dạng nào
// Time là thời điểm nhận được tin nhắn của client từ server
type RoomData struct {
	UserId uuid.UUID
	Data   []byte
	Time   int64
}

// must be implemented
// lớp handler xử lý các logic trong phòng chơi
// với mỗi game phải implement interface này để có thể lấy dữ liệu từ phòng
type RoomHandler interface {
	// reload phòng từ data lưu trong cache sau khi khởi động lại server
	Reload(room Room, state string) interface{}
	// khởi tạo phòng
	Created(room Room) interface{}
	// kiểm tra điều kiện tham gia phòng của người chơi
	AllowJoin(room Room, userId uuid.UUID) *CheckJoinConditionResult
	// một action được schedule trước đó sau khi đến lịch sẽ nhảy vào hàm này
	Processor(room Room, action string, data map[string]interface{}) interface{}
	// người chơi tham gia vào phòng
	OnJoined(room Room, userId uuid.UUID) interface{}
	// xử lý request yêu cầu thoát của người chơi
	OnExitRequest(room Room, userId uuid.UUID)
	// người chơi thoát khỏi phòng hoặc bị kick
	OnLeft(room Room, userId uuid.UUID) interface{}
	// người chơi mất kết nối với server
	OnDisconnected(room Room, userId uuid.UUID) interface{}
	// mỗi tick rate sẽ nhảy vào game loop một lần
	// xử lý tất cả các request của người chơi
	Loop(room Room, roomChan <-chan *RoomData) interface{}
	// phòng đã được đóng lại
	OnClose(room Room)
}

// tạo mới một room handler
type CreateRoomHandlerFunc func() RoomHandler

func RegisterHandler(gameId string, f CreateRoomHandlerFunc) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	if f == nil {
		panic("game api: register handler func is nil")
	}
	if _, dup := handlers[gameId]; dup {
		panic("game api: register called twice for game id " + gameId)
	}
	handlers[gameId] = f
}

func GetHandlers() map[string]CreateRoomHandlerFunc {
	return handlers
}

func GetHandler(gameId string) CreateRoomHandlerFunc {
	return handlers[gameId]
}