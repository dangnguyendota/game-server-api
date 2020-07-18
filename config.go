package gsi

import "github.com/gocql/gocql"

type ConfigAll struct {
	Pusher          *PusherConfig
	Database        *DatabaseConfig
	WebsocketServer *WebsocketServerConfig
	MatchingServer  *MatchingServerConfig
}

type PusherConfig struct {
	RestApiKey string
	AppID      string
	ApiUrl     string
	Workers    int
	PoolQueue  int
}

type DatabaseConfig struct {
	Hosts         []string
	KeySpace      string
	Consistency   gocql.Consistency
	MemcachedHost string
}

type WebsocketServerConfig struct {
	JwtSecret                  string
	TokenExpiredTime           int
	MaxMessageLength           int
	MaxCCU                     int
	MaxClientMessageQueue      int
	MaxClientMessageLength     int
	JoinLeaveEventChanSize     int
	RoomEventChanSize          int
	RoomDataEventSize          int
	RoomCheckConditionChanSize int
	DirectorLoopTime           int64
}

type MatchingServerConfig struct {
	N int
}
