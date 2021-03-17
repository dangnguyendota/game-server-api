package api

import "github.com/google/uuid"

// Mỗi Room đều chứa 1 Bot dùng để làm việc với Bot được viết sẵn trước đó
//  Bot là các người chơi ảo dùng để giả làm người trong game
type Bot interface {
	// hàm Get gọi Bot vào trong bàn
	Get(callback func(err error))
	// hàm Call sẽ gửi 1 request lên Bot server
	// sau khi Bot server nhận được request sẽ xử lý request đó và trả về
	// một response trong hàm callback
	// returns error nếu call Bot thất bại
	Call(botId uuid.UUID, request string, callback func(response string)) error
}
