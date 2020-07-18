package gsi

import "github.com/gocql/gocql"

type ConfigAll struct {
	Pusher         *PusherConfig
	Database       *DatabaseConfig
	Matchmaker     *MatchmakerConfig
	MatchingServer *MatchingServerConfig
	ApiServer      *ApiServerConfig
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

type MatchmakerConfig struct {
	Host                       string
	Port                       string
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
	Host string
	Port string
	N    int
}

type ApiServerConfig struct {
	MetricsHost     string
	MetricsPort     string
	LeaderBoardHost string
	LeaderBoardPort string
	AdminHost       string
	AdminPort       string
	BotHost         string
	BotPort         string
}
