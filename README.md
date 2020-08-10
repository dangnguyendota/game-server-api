# game server api

* **Room Handler** _(must implement)_**:** override lại để thêm xử lý logic cho game.
* **Room** _(fixed)_**:** chứa thông tin của phòng.
* **Notification Pusher** _(fixed)_**:** bắn thông báo qua one signal
* **Scheduler** _(fixed)_**:** lên lịch cho các hành động, khởi tạo riêng cho mỗi room một scheduler.
* **User:** đại điện cho một phiên đăng nhập của người chơi bao gồm Sid là session id và Id là user id.
* **Mode:** hai mode game authoritative server cho modify logic game và relayed server chuyển tiếp gói tin thời gian thực không qua xử lý.
* **Database** _(must implement)_**:** chứa các hàm xử lý các request trong database.