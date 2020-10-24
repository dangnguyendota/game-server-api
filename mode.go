package api

// Mode game của một Room.
// có hai loại game chính là turn based và realtime.
// ứng với hai loại thường là hai mode game chính đó là
// AuthoritativeMultiPlayerServer cho dòng game turn based
// hoặc RelayedMultiPlayerServer cho dòng game realtime.
// đối với AuthoritativeMultiPlayerServer yêu cầu Room phải được import một RoomHandler
// giúp xử lý các request từ người chơi hoặc các action của server.
// đối với RelayedMultiPlayerServer thì server sẽ tự động gom các tin nhắn trong
// hàng đợi vào với nhau và gửi lại cho tất cả mọi người trong Room sau mỗi một Room.TickRate.
type Mode uint8

const (
	AuthoritativeMultiPlayerServer = 0
	RelayedMultiPlayerServer       = 1
)
