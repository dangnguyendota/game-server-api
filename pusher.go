package api

// NotificationPusher giúp bắn thông báo về máy người chơi thông qua dịch vụ
// của bên thứ 3.
type NotificationPusher interface {
	// bắn thông báo cho người chơi
	// message là thông báo muốn gửi về client.
	// f là hàm xử lý response sẽ được gọi đến sau khi request gửi đi thành công.
	Push(message *NotificationMessage, f func(response []byte))
	// dừng NotificationPusher lại và xóa tất cả các con goroutines
	Stop()
}

type NotificationMessage struct {
	AppID            string            `json:"app_id"`
	IncludePlayerIds []string          `json:"include_player_ids"`
	IncludedSegments []string          `json:"included_segments"`
	Data             map[string]string `json:"data"`
	Contents         map[string]string `json:"contents"`
	Headings         map[string]string `json:"headings"`
	Url              string            `json:"url"`
	BigPicture       string            `json:"big_picture"`
	LargeIcon        string            `json:"large_icon"`
}

func NewNotificationMessage(header, content string) *NotificationMessage {
	return &NotificationMessage{
		Headings: map[string]string{"en": header},
		Contents: map[string]string{"en": content},
	}
}

func (n *NotificationMessage) AddPlayer(id string) *NotificationMessage {
	n.IncludedSegments = []string{}
	if n.IncludePlayerIds == nil {
		n.IncludePlayerIds = []string{id}
	} else {
		n.IncludePlayerIds = append(n.IncludePlayerIds, id)
	}
	return n
}

func (n *NotificationMessage) AddAllPlayers() *NotificationMessage {
	n.IncludePlayerIds = []string{}
	n.IncludedSegments = []string{"All"}
	return n
}

func (n *NotificationMessage) WithUrl(url string) *NotificationMessage {
	n.Url = url
	return n
}

func (n *NotificationMessage) WithBigPicture(url string) *NotificationMessage {
	n.BigPicture = url
	return n
}

func (n *NotificationMessage) WithAndroidIcon(url string) *NotificationMessage {
	n.LargeIcon = url
	return n
}

func (n *NotificationMessage) AddData(key, value string) *NotificationMessage {
	if n.Data == nil {
		n.Data = map[string]string{key: value}
	} else {
		n.Data[key] = value
	}
	return n
}
