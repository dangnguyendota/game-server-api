package api

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// dữ liệu lưu trữ trong một Room từ khi khởi tạo
// dữ liệu này là conf không thể thay đổi nhưng lấy ra được interface Room
// giúp lưu trữ các conf trong game lúc khởi tạo phòng trong RoomHandler
type Metadata interface {
	// convert sang string
	String() string
	// convert sang byte
	Byte() (byte, error)
	// convert sang byte array
	Bytes() []byte
	// convert sang int64
	Int64() (int64, error)
	// convert sang int32
	Int32() (int32, error)
	// convert sang uint64
	Uint64() (uint64, error)
	// convert sang uint32
	Uint32() (uint32, error)
	// convert sang float64
	Float64() (float64, error)
	// convert sang float32
	Float32() (float32, error)
	// parse Metadata sang một struct
	Json(interface{}) error
}

// mỗi Room chứa dữ liệu của một phòng
// nó chứa một cụm người chơi, quản lý một số đối tượng
type Room interface {
	// mỗi phòng có một ID duy nhất dưới dạng uuid
	ID() uuid.UUID
	// mỗi game sẽ được phân biệt bằng một game id duy nhất
	// nó được nhà phát triển quy định.
	// Game trả về game id của game ứng với Room này
	Game() string
	// Node trả về tên của cụm máy chủ chứa những người chơi trong phòng này.
	// mỗi phòng sẽ được tạo từ một request của máy chủ với tên Node.
	// phòng không thể tạo trực tiếp từ client mà phải thông qua một Server khác.
	Node() string
	// trả về Mode server ứng với Room này
	//// có hai loại Mode là
	// Authoritative MultiPlayer Server ứng với loại game cần control logic
	// và Relayed MultiPlayer Server ứng với loại game real-time hoặc không cần control logic.
	Mode() Mode
	// fps tối đa của game, với mỗi giây, server sẽ xử lý request
	// tối đa TickRate lần.
	TickRate() int64
	// MaxPlayers trả về số lượng tối đa người chơi có trong phòng.
	MaxPlayers() int
	// MinPlayers trả về số lượng tối thiểu người chơi có trong phòng.
	MinPlayers() int
	// MaxViewers trả về số lượng tối đa người xem vào được phòng.
	MaxViewers() int
	// Logger giúp ghi log là một zap.Logger
	Logger() *zap.Logger
	// trả về dữ liệu được config sẫn trong Room
	Metadata(name string) Metadata
	// trả về trạng thái của game, nếu State = nil thì phòng sẽ tự động hủy.
	// trạng thái sẽ được update sau mỗi hàm trong RoomHandler được khởi tạo.
	State() interface{}
	// trả về danh sách người chơi đang online có trong phòng.
	Players() []uuid.UUID
	// trả về danh sách người xem đang online có trong phòng.
	Viewers() []uuid.UUID
	// trả về thời gian khởi tạo của phòng.
	CreateTime() int64
	// trả về thời gian mà lần cuối Room được update State
	LastUpdateTime() int64
	// đuổi một người chơi (hoặc người xem) ra khỏi Room.
	Kick(userId uuid.UUID) error
	// hủy phòng này, Kick tất cả người chơi ra khỏi Room.
	Destroy()
	// tạm dừng trò chơi, tất cả mọi request đều sẽ bị Room reject và
	// hàm Loop trong RoomHandler sẽ không được gọi đến nữa.
	Pause()
	// tiếp tục trò chơi.
	Unpause()
	// trả về true nếu phòng đang bị tạm dừng.
	Paused() bool
	// Scheduler giúp lên lịch cho các hành động tiếp theo của server.
	// nó rất hữu ích trong các trường hợp cần lên lịch cho một hành động sau
	// một khoảng thời gian nhất định một cách tự động.
	Scheduler() Scheduler
	// một Pusher cung cấp các functions giúp gửi thông báo về máy người chơi.
	// có thể hữu ích trong một số trường hợp muốn thông báo cho người chơi
	// kể cả khi người chơi đã tắt game.
	Pusher() NotificationPusher
	// chứa DB giúp làm việc với database.
	// lớp DB sẽ được tạo mới tùy ứng với từng game.
	// đễ sử dụng DB thì phải ép kiểu sang dạng mà mình đã import vào.
	DB() Database
	// chứa các phương thức làm việc với Bot
	Bot() Bot
	// trả về Dispatcher của Room tương ứng.
	Dispatcher() Dispatcher
}
