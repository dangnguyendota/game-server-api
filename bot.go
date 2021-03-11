package api

import "github.com/google/uuid"

// Mỗi Room đều chứa 1 Bot dùng để làm việc với Bot được viết sẵn trước đó
//  Bot là các người chơi ảo dùng để giả làm người trong game
type Bot interface {
	// hàm Get gọi Bot vào trong bàn
	// nếu gọi Bot thành công sẽ trả về callback chứa ID của Bot
	// returns error nếu gọi Bot thất bại
	Get(callback func(botId uuid.UUID)) error
	// hàm Remove xóa (đuổi) Bot khỏi Room
	// nếu xóa Bot thành công thì sẽ trả về 1 callback chứa ID của Bot bị kick
	// returns error nếu xóa Bot thất bại
	Remove(botId uuid.UUID, callback func(botId uuid.UUID)) error
	// hàm Call sẽ gửi 1 request lên Bot server
	// sau khi Bot server nhận được request sẽ xử lý request đó và trả về
	// một response trong hàm callback
	// returns error nếu call Bot thất bại
	Call(request string, callback func(response string)) error
}
