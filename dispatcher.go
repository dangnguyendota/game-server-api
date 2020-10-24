package api

import "github.com/google/uuid"

// message router giúp gửi tin nhắn cho người chơi đang ở trong một phòng
type Dispatcher interface {
	// gửi tin nhắn đến một người chơi
	Send(userId uuid.UUID, data []byte)
	// gửi tin nhắn đến tất cả mọi người trong phòng
	SendAll(data []byte)
	// gửi tin nhắn cho tất cả mọi người trừ một người
	// sẽ rất hữu ích cho các trường hợp thông báo cho mọi người khác
	// về sự thay đổi trạng thái của người chơi nào đó mà không cần thông báo
	// cho người chơi đó
	SendAllExceptOne(userId uuid.UUID, data []byte)
	// gửi cho tất cả mọi người đang chơi
	SendAllPlayers(data []byte)
	// tương tự SendAllExceptOne nhưng chỉ gửi cho những người đang chơi
	SendAllPlayersExceptOne(userId uuid.UUID, data []byte)
	// gửi cho tất cả mọi người đang xem
	SendAllViewers(data []byte)
	// tương tự SendAllExceptOne nhưng chỉ gửi cho những người đang xem.
	SendAllViewersExceptOne(userId uuid.UUID, data []byte)
	// gửi một tin nhắn thông báo lỗi đến người chơi
	SendError(userId uuid.UUID, code uint32, message string)
}
