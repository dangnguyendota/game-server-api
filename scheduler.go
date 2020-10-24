package api

import "time"

// mỗi Scheduler cung cấp các hàm để lên lịch hoặc hủy lịch cho các hành động mà
// server muốn thực hiện.
type Scheduler interface {
	// Schedule một action.
	// action là tên của action đó, tên này là duy nhất, nếu một action mới
	// trùng tên với action trước đó đã được Schedule mà vẫn đang đợi thì sẽ trả
	// về một error và action đó không được thực hiện.
	// data là dữ liệu được truyền theo, khi action này được thực hiện có thể lấy data ra.
	// d là thời gian chờ đến khi mà action sẽ được thực hiện.
	Schedule(action string, data map[string]interface{}, d time.Duration) error
	// hủy một hành động đang được Schedule, nếu hành động đó không tồn tại hoặc
	// đã được thực hiện rồi thì sẽ trả về một error
	Cancel(action string) error
	// kiểm tra xem action có đang trong hàng đợi không.
	Waiting(action string) bool
	// hủy tất cả các action đang trong hàng đợi.
	CancelAll()
	// dừng Scheduler, xóa tất cả các con goroutine của Scheduler.
	Stop()
}